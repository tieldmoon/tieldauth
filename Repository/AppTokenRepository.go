package Repository

import (
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
	return false
}
