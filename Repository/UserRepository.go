package Repository

import (
	"context"

	"github.com/tieldmoon/tieldauth/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByEmail(email string) (Models.User, bool)
}

type UserRepositoryMongo struct {
	Client *mongo.Client
}

func (u *UserRepositoryMongo) FindByEmail(email string) (Models.User, bool) {
	col := u.Client.Database("User").Collection("user_data")
	var result bson.D
	col.FindOne(context.TODO(), bson.D{
		bson.E{Key: "email", Value: email},
	}).Decode(&result)
	doc, err := bson.Marshal(result)
	if err != nil {
		return Models.User{}, false
	}
	var data Models.User
	err = bson.Unmarshal(doc, &data)
	if err != nil {
		return Models.User{}, false
	}
	return data, true

}
