package supplier

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	// 待受を開始する
	// イベントが来る
	// debounce, throttleする
	//   情報を検索する
	//   解析情報が新規モノであれば投入する
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
			if !initial {
				log.Printf("data: %v", result)
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
