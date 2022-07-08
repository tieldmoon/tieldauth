package Service

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnPool() *mongo.Client {
	mo := make(chan *mongo.Client)
	go func(mo chan *mongo.Client) {
		m, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_SERVER")))
		if err != nil {
			fmt.Println(err)
		}
		initMongoDb(m)
		log.Println("Success connect mongodb")
		mo <- m
	}(mo)
	mongodb := <-mo
	close(mo)
	return mongodb
}

func initMongoDb(c *mongo.Client) {
	// run if not exits, is exits skip
	if err := c.Database("Token").CreateCollection(context.TODO(), "app_token"); err == nil {
		// generate sample data for local development
		c.Database("Token").Collection("app_token").InsertOne(
			context.TODO(), bson.D{
				bson.E{Key: "app_id", Value: "17238916"},
				bson.E{Key: "app_key", Value: "dhw92ujie982ujdiq982uehd1d2"},
				bson.E{Key: "publisher", Value: "test@samplemail.com"},
			},
		)
	}
}
