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
	AppPort  int
	Postgres postgres
}

func NewDefault() *AppConfig {
	return &AppConfig{
		AppPort: 8080,
		Postgres: postgres{
			Host:     "db",
			Name:     "shortener",
			User:     "shortener",
			Password: "qwerty",
			SSLMode:  "disable",
			Port:     5432,
		},
	}
}
