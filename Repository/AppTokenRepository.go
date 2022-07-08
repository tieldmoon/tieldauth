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

func (t *TokenRepositoryMongo) CheckAppIdIsAvailable(app_id string) (Models.AppToken, bool) {
	col := t.Client.Database("Token").Collection("app_token")
	var result bson.D
	col.FindOne(context.TODO(), bson.D{
		bson.E{Key: "app_id", Value: app_id},
	}).Decode(&result)
	var data Models.AppToken
	doc, err := bson.Marshal(result)
	if err != nil {
		fmt.Println("er 1:", err)
	}
	err = bson.Unmarshal(doc, &data)
	if err != nil {
		fmt.Println("err 2:", err)
	}

	return data, false
}
