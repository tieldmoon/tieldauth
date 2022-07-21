package Models

type User struct {
	Id        string `bson:"id"`
	Email     string `bson:"email"`
	Username  string `bson:"username"`
	Password  string `bson:"password"`
	Birthdate string `bson:"birthdate"`
}
