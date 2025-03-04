package config

type Config struct {
	Port   int
	NATURL string
	DBURL  string
}

func GetConfig() Config {
	return Config{
		Port:   4987,
		NATURL: "nats://localhost:4222",
	}
}
