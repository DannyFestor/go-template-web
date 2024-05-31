package config

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: ":4000",
	}
}
