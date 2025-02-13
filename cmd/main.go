package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/logger"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/pskoob/puffmi.git/config"
	"go.uber.org/zap"
)

var cfg config.con

func main() {
	envconfig.MustProcess("", &config)

	if err := logger.BuildLogger(config.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}

	zap.L().Info("PostgresURI: ", zap.String("uri", config.PostgresURI))
	zap.L().Info("MigrationsDir: ", zap.String("dir", config.MigrationsDir))

	time.Sleep(3 * time.Second)

	migrator := postgresql.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	defer sqlxConn.Close()

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqalxConn.Close()

	zap.L().Info("Database manage was process successfully")
}

func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь можно добавить хеширование пароля перед сохранением
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User  registered successfully", "user": user})
}
