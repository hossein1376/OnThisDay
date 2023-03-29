package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/slog"
)

var RDB *redis.Client

type config struct {
	port string
	db   struct {
		address  string
		password string
		db       int
	}
}

type application struct {
	logger *slog.Logger
	config *config
}

func RunServer() {
	jsonHandler := slog.NewJSONHandler(os.Stdout)
	logger := slog.New(jsonHandler)

	cfg := &config{
		port: "4000",
		db: struct {
			address  string
			password string
			db       int
		}{
			address:  "localhost:6379",
			password: "",
			db:       0,
		},
	}

	app := &application{
		logger: logger,
		config: cfg,
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.db.address,
		Password: cfg.db.password,
		DB:       cfg.db.db,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		app.logger.Error("Redis connection failed", slog.String("error:", err.Error()))
		return
	}

	app.logger.Info(fmt.Sprintf("Starting server on port %v..", cfg.port))
	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), app.router())
	if err != nil {
		app.logger.Error("error starting the server", slog.String("error:", err.Error()))
		return
	}
}
