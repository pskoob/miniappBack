package config

type Config struct {
	Addr           string `envconfig:"ADDR"`
	MigrationsDir  string `envconfig:"MIGRATIONS_DIR"`
	PostgresURI    string `envconfig:"POSTGRES_URI"`
	LogLevel       string `envconfig:"LOG_LEVEL"`
	ApiKey         string `envconfig:"API_KEY"`
	SecretKey      string `envconfig:"SECRET_KEY"`
	SendingAccount string `envconfig:"SENDING_ACCOUNT"`
	TokenSecretKey string `envconfig:"TOKEN_SECRET_KEY"`
}
