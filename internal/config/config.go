package config

type Config struct {
	NATURL string
	DBURL  string
}

func GetConfig(natURL string, dbURL string) Config {
	return Config{
		NATURL: "nats://localhost:4222",
		DBURL:  dbURL,
	}
}
