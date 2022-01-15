package config

type AppConfig struct {
	Port          string         `json:"port"`
	MongoSettings *MongoSettings `json:"mongo"`
}

type MongoSettings struct {
	ConnectionURI string `json:"connectionURI"`
	Database      string `json:"database"`
	Collection    string `json:"collection"`
}
