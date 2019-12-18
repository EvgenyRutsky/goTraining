package client

type Config struct {
	ClientURI string
	Dbname string
	BooksCollection string
	ReadersCollection string
}

func NewConfig() *Config {
	return &Config{
		ClientURI: "mongodb://127.0.0.1:27017",
		Dbname:    "dev",
		BooksCollection:"books",
		ReadersCollection:"readers",
	}
}
