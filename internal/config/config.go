package config

type postgres struct {
	Host     string
	Name     string
	User     string
	Password string
	SSLMode  string
	Port     int
}

type AppConfig struct {
	HTTPPort int
	RPCPort  int
	Postgres postgres
}

func NewDefault() *AppConfig {
	return &AppConfig{
		RPCPort:  3003,
		HTTPPort: 8080,
		Postgres: postgres{
			Host:     "localhost",
			Name:     "shortener",
			User:     "shortener",
			Password: "qwerty",
			SSLMode:  "disable",
			Port:     5436,
		},
	}
}
