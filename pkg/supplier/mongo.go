package supplier

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
	"github.com/p2pquake/userquake-aggregator/pkg/evaluate"
	"github.com/p2pquake/userquake-aggregator/pkg/rate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const epspTimeFormat = "2006/01/02 15:04:05.000"

var loc, _ = time.LoadLocation("Asia/Tokyo")

type Mongo struct {
	Done chan struct{}

	context    context.Context
	client     *mongo.Client
	collection *mongo.Collection
}

func (m *Mongo) Start(context context.Context, uri string, database string, collection string) {
	m.Done = make(chan struct{}, 1)
	m.context = context

	go m.run(uri, database, collection)
}

func (m *Mongo) run(uri string, db string, c string) {
	defer func() { m.Done <- struct{}{} }()

	op := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(op)
	if err != nil {
		log.Fatal(err)
	}
	m.client = client

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	col := client.Database(db).Collection(c)
	m.collection = col

	ct := options.TailableAwait
	wait := time.Duration(1)
	options := options.FindOptions{CursorType: &ct, MaxAwaitTime: &wait}
	filter := bson.D{}
	cur, err := col.Find(context.Background(), filter, &options)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	limiter := rate.NewLimiter(1, m.calc)
	initial := true

L:
	for {
		select {
		case <-m.context.Done():
			log.Printf("Stopping tail due to context cancellation")
			break L
		default:
		}

		if cur.TryNext(context.TODO()) {
			var result bson.M
			if err := cur.Decode(&result); err != nil {
				log.Fatal(err)
			}
			if !initial && result["code"].(int32) == 561 {
				limiter.Run()
			}
			continue
		} else {
			initial = false
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		if cur.ID() == 0 {
			break
		}
	}
}

func (m *Mongo) calc() {
	beginTime := time.Now().Add(-30 * time.Minute).In(loc).Format("2006/01/02 15:04:05.000")

	filters := bson.M{"code": bson.M{"$in": []int{555, 561}}, "time": bson.M{"$gte": beginTime}}
	opt := options.FindOptions{Sort: bson.D{{"time", int64(1)}}}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cur, err := m.collection.Find(ctx, filters, &opt)
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	items := make([]bson.M, 0)
	cur.All(ctx, &items)

	// aggregate & evaluate
	body, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("JSON marshal error: %v from %v", err, items)
	}

	// parse
	epspRecords := []epsp.Record{}
	err = json.Unmarshal(body, &epspRecords)
	if err != nil {
		log.Fatalf("JSON unmarshal error: %v from %v", err, string(body))
	}

	// aggregate & evaluate
	aggregationResults := aggregate.CompatibleAggregator{}.Aggregate(epspRecords)

	evaluationResults := []evaluate.Result{}

	for _, r := range aggregationResults {
		result := evaluate.CompatibleEvaluator{}.Evaluate(r)
		evaluationResults = append(evaluationResults, result)
	}

	for _, r := range evaluationResults {
		// 存在しなければ足す.
		filters := bson.M{"code": 9611, "started_at": r.StartedAt.Format(epspTimeFormat), "count": r.Count}

		count, err := m.collection.CountDocuments(context.Background(), filters)
		if err != nil {
			log.Fatalf("Count error: %v", err)
		}

		if count > 0 {
			continue
		}

		log.Printf("insert: %v", convert(r))

		if _, err = m.collection.InsertOne(context.Background(), convert(r)); err != nil {
			log.Fatalf("Insert error: %v", err)
		}
	}
}

func convert(r evaluate.Result) map[string]interface{} {
	return map[string]interface{}{
		"code":             9611,
		"time":             time.Now().In(loc).Format(epspTimeFormat),
		"started_at":       r.StartedAt.Format(epspTimeFormat),
		"updated_at":       r.UpdatedAt.Format(epspTimeFormat),
		"count":            r.Count,
		"confidence":       r.Confidence,
		"area_confidences": convertAreaConfidence(r.AreaConfidence),
	}
}

func convertAreaConfidence(a map[epsp.AreaCode]evaluate.AreaResult) map[int]map[string]interface{} {
	r := map[int]map[string]interface{}{}
	for k, v := range a {
		r[int(k)] = map[string]interface{}{"confidence": float64(v.Confidence), "count": v.Count, "display": v.Display()}
	}
	return r
}
