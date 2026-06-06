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
	Done   chan struct{}
	DryRun bool // true なら InsertOne を行わずログ出力のみ

	context    context.Context
	client     *mongo.Client
	collection *mongo.Collection
	limiter    *rate.Limiter
	evalCache  map[evalCacheKey]evaluate.Result
}

// evalCacheKey は1クラスタの評価結果メモ化キー。
// 同じ StartedAt かつ件数が同じならクラスタ内容も同一なので、Evaluate を再実行せず使い回せる。
type evalCacheKey struct {
	startedAt string
	count     int
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

	connectCtx, connectCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer connectCancel()

	err = client.Connect(connectCtx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = client.Disconnect(ctx)
	}()

	col := client.Database(db).Collection(c)
	m.collection = col

	ct := options.TailableAwait
	wait := 500 * time.Millisecond
	findOpts := options.FindOptions{CursorType: &ct, MaxAwaitTime: &wait}
	filter := bson.D{}
	cur, err := col.Find(context.Background(), filter, &findOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = cur.Close(ctx)
	}()

	m.evalCache = map[evalCacheKey]evaluate.Result{}
	m.limiter = rate.NewLimiter(1, m.calc)
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
				m.limiter.Run()
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
	opt := options.FindOptions{Sort: bson.M{"time": int64(1)}}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cur, err := m.collection.Find(ctx, filters, &opt)
	if err != nil {
		log.Fatalf("Find error: %v", err)
	}
	defer cur.Close(ctx)

	items := make([]bson.M, 0)
	cur.All(ctx, &items)

	log.Printf("Aggregate from %v items", len(items))

	// to json
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

	// aggregate & evaluate（クラスタ単位でメモ化）
	evaluationResults := m.aggregateAndEvaluateCached(epspRecords)

	// 適応スロットル: 最大クラスタ件数に応じて次回以降の calc 間隔を調整
	maxCount := 0
	for _, r := range evaluationResults {
		if r.Count > maxCount {
			maxCount = r.Count
		}
	}
	m.limiter.SetInterval(adaptiveInterval(maxCount))

	for _, r := range evaluationResults {
		// 存在しなければ足す.
		filters := bson.M{"code": 9611, "started_at": r.StartedAt.Format(epspTimeFormat), "count": r.Count}

		countCtx, countCancel := context.WithTimeout(context.Background(), 10*time.Second)
		count, err := m.collection.CountDocuments(countCtx, filters)
		countCancel()
		if err != nil {
			log.Fatalf("Count error: %v", err)
		}

		if m.DryRun {
			body, _ := json.Marshal(convert(r))
			if count > 0 {
				log.Printf("[dry-run] skip (already exists): %s", body)
			} else {
				log.Printf("[dry-run] would insert: %s", body)
			}
			continue
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
		"user-agent":       "userquake-aggregator 20260610",
		"ver":              "20260610",
	}
}

func convertAreaConfidence(a map[epsp.AreaCode]evaluate.AreaResult) map[int]map[string]interface{} {
	r := map[int]map[string]interface{}{}
	for k, v := range a {
		r[int(k)] = map[string]interface{}{"confidence": float64(v.Confidence), "count": v.Count, "display": v.Display()}
	}
	return r
}

// aggregateAndEvaluateCached はクラスタ単位で Evaluate 結果をキャッシュする。
// 30 分窓から外れたクラスタはこの呼び出しで再構築される newCache に含まれないため、自然に落ちる。
func (m *Mongo) aggregateAndEvaluateCached(records []epsp.Record) []evaluate.Result {
	aggregationResults := aggregate.CompatibleAggregator{}.Aggregate(records)

	newCache := map[evalCacheKey]evaluate.Result{}
	evaluationResults := []evaluate.Result{}

	for _, ar := range aggregationResults {
		if len(ar.Userquakes) == 0 {
			continue
		}
		key := evalCacheKey{
			startedAt: ar.StartedAt.Time.Format(epspTimeFormat),
			count:     len(ar.Userquakes),
		}

		result, ok := m.evalCache[key]
		if !ok {
			result = evaluate.CompatibleEvaluator{}.Evaluate(ar)
		}
		newCache[key] = result

		if time.Since(*result.StartedAt.Time) > 25*time.Minute {
			continue
		}
		evaluationResults = append(evaluationResults, result)
	}

	m.evalCache = newCache

	return evaluationResults
}

// adaptiveInterval は最大クラスタ件数に応じて、次回以降の calc 間隔（秒）を返す。
// 整数秒。地震時の高負荷を吸収しつつ、平常時の反応性を保つ。
func adaptiveInterval(maxCount int) float64 {
	switch {
	case maxCount >= 200:
		return 3
	case maxCount >= 50:
		return 2
	default:
		return 1
	}
}
