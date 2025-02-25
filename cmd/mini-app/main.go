package main

import (
	// "context"
	// "fmt"
	"log"
	"miniappBack/pkg/logger"
	"miniappBack/internal/handler"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"miniappBack/config"
	"miniappBack/database/postgresql"

	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"

	"github.com/justinas/alice"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var (
	cfg config.Config
)make generate

func main() {
	envconfig.MustProcess("", &cfg)

	if err := logger.BuildLogger(cfg.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}

	zap.L().Info("PostgresURI: ", zap.String("uri", cfg.PostgresURI))
	zap.L().Info("MigrationsDir: ", zap.String("dir", cfg.MigrationsDir))

	time.Sleep(3 * time.Second)

	migrator := postgresql.NewMigrator(cfg.PostgresURI, cfg.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", cfg.PostgresURI)
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

		appHandler := handler.New()

		chain := alice.New(appHandler.WsMiddleware).Then(appHandler)
		if chain == nil {
			fmt.Println(chain)
		}

		server := http.Server{
			Handler: chain,
			Addr:    ":" + cfg.Addr,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Блокируем `main`, ожидая сигнала
		<-quit
		zap.L().Info("Shutdown server ...")

		// Создаем таймаут для выключения сервера
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Завершение сервера
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown:", err)
		}
		zap.L().Info("Server exiting")

	}

