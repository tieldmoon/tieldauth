package Service

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Worker struct {
	Wg    *sync.WaitGroup
	Jobs  chan map[int]any
	Mongo chan *mongo.Client
}

func (w *Worker) InitWorker(worker_number int) {
	fmt.Println(os.Getenv("MONGO_SERVER"))
	con, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_SERVER")))
	if err != nil {
		panic(err)
	}
	initMongoDb(con)
	w.Mongo <- con
	for i := 0; i < worker_number; i++ {
		go func(i int, w *Worker) {

		}(i, w)
	}
}

func initMongoDb(c *mongo.Client) {
	// run if not exits, is exits skip
	if err := c.Database("Token").CreateCollection(context.TODO(), "app_token"); err == nil {
		// generate sample data
		c.Database("Token").Collection("app_token").InsertOne(
			context.TODO(), bson.D{
				bson.E{Key: "app_id", Value: "17238916"},
				bson.E{Key: "app_key", Value: "dhw92ujie982ujdiq982uehd1d2"},
				bson.E{Key: "publisher", Value: "tieldmoon"},
			},
		)
	}
}
