package repository

type TokenRepository interface {
	CheckAppIdIsAvailable(app_id string) bool
	GetAppKey(app_id string)
	VerifySecretToken(secret_key string)
}

type TokenRepositoryMongo struct {
}

func (t *TokenRepositoryMongo) CheckAppIdIsAvailable(app_id string) bool {

	return false
}
