package config

type Config struct {
	Addr          string `envconfig:"ADDR" default:"8000"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"../../database/postgresql/migrations"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://pasha:pasha@localhost:5449/mini_app_db?sslmode=disable"`
	LogLevel      string `envconfig:"LOG_LEVEL"`
	ApiKey        string `envconfig:"API_KEY" default:"8071867487:AAH1c8_uLZOXqLUgAJ9orKGQRtjc04y3Okk"`
}
