package store

type Config struct {
	DB string
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DB: "mongodb",
		DatabaseURL: "mongodb://127.0.0.1:27017",
	}
}