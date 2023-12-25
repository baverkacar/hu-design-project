package repository

type UserRepository interface {
	Get() string
	ConnectToMongo(mongoURL string) error
}
