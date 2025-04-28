package config

import (
	"time"
)

type Config struct {
	Addr                     string        `envconfig:"ADDR"`
	MigrationsDir            string        `envconfig:"MIGRATIONS_DIR"`
	PostgresURI              string        `envconfig:"POSTGRES_URI"`
	LogLevel                 string        `envconfig:"LOG_LEVEL"`
	ApiKey                   string        `envconfig:"API_KEY"`
	SecretKey                string        `envconfig:"SECRET_KEY"`
	SendingAccount           string        `envconfig:"SENDING_ACCOUNT"`
	TokenSecretKey           string        `envconfig:"TOKEN_SECRET_KEY"`
	AutoClickerWorkTime      time.Duration `envconfig:"AUTO_CLICKER_WORK_TIME"`
	AutoClickerTickTime      time.Duration `envconfig:"AUTO_CLICKER_TICK_TIME"`
	AutoClickerSaveTicksTime time.Duration `envconfig:"AUTO_CLICKER_SAVE_TICKS_TIME"`
	EnergyBaseCapacity       int64         `envconfig:"ENERGY_BASE_CAPACITY"`
	EnergyUpgrade            int64         `envconfig:"ENERGY_UPGRADE"`
	EnergyTickTime           time.Duration `envconfig:"ENERGY_TICK_TIME"`
	EnergySaveTicksTime      time.Duration `envconfig:"ENERGY_SAVE_TICKS_TIME"`
	BotLink                  string        `envconfig:"BOT_LINK"`
	HashSalt                 string        `envconfig:"HASH_SALT"`
	ReferralUserPart         float64       `envconfig:"REFERRAL_USER_PART"`
}
