package Repository

type UserRepository interface {
	Auth(u string, p string) (bool, error)
}
