package api

import (
	"context"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func RunServer() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on port 4000..\n")
	err = http.ListenAndServe(":4000", router())
	if err != nil {
		log.Fatal(err)
	}
}
