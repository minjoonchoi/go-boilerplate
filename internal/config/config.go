package config

type Config struct {
	Server struct {
		Address string
	}
}

func Load() (*Config, error) {
	return &Config{
		Server: struct{ Address string }{Address: ":8080"},
	}, nil
} 