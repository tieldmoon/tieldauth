package Models

type AppToken struct {
	Appid     string `bson:"app_id"`
	AppKey    string `bson:"app_key"`
	Publisher string `bson:"publisher"`
}
