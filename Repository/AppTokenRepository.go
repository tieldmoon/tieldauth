package Repository

import (
	"context"
	"fmt"

	"github.com/tieldmoon/tieldauth/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	CheckAppIdIsAvailable(app_id string) bool
	GetAppKey(app_id string) string
}

type TokenRepositoryMongo struct {
	Client *mongo.Client
}

func (t *TokenRepositoryMongo) CheckAppIdIsAvailable(app_id string) bool {
	col := t.Client.Database("Token").Collection("app_token")
	sr := col.FindOne(context.TODO(), bson.D{
		bson.E{Key: "app_id", Value: app_id},
	})
	var result bson.D
	var data Models.AppToken
	sr.Decode(&result)
	doc, err := bson.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	err = bson.Unmarshal(doc, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	return false
}
