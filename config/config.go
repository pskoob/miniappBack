package config

var config struct {
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"../../internal/database/postgresql/migrations"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://postgres:yourpassword@postgres:5432/yourdbname?sslmode=disable"`
}
