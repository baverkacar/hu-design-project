package configuration

// MongoConfig veritabanı yapılandırmalarını saklamak için bir struct.
type MongoConfig struct {
	URL                 string
	Database            string
	UserCollection      string
	DevicesCollection   string
	RequestsCollection  string
	AlertsCollection    string
	BlacklistCollection string
	WhitelistCollection string
}

// NewMongoConfig, yeni bir MongoConfig nesnesi oluşturur ve döndürür.
func NewMongoConfig() *MongoConfig {
	return &MongoConfig{
		URL:                 "mongodb://localhost:27017",
		Database:            "hu-design-project",
		UserCollection:      "user",
		DevicesCollection:   "devices",
		RequestsCollection:  "requests",
		AlertsCollection:    "alerts",
		BlacklistCollection: "blacklist",
		WhitelistCollection: "whitelist",
	}
}
