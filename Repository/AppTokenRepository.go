package Repository

import (
	"context"

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
	col.Find(context.TODO(), bson.D{
		bson.E{Key: "app_id", Value: app_id}
	})
	return false
}
