package config

type Config struct {
	Addr          string `envconfig:"ADDR" default:"8000"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"../../miniappBack/database/postgresql/migrations"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://yourusername:yourpassword@localhost:5432/yourdbname?sslmode=disable"`
	LogLevel      string `envconfig:"LOG_LEVEL"`
}
