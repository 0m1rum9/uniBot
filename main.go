package main

import (
	"context"
	"fmt"
	"time"

	"os"
	DB "unibot/internal/db"
	"unibot/internal/handlers"
	"unibot/internal/services"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gopkg.in/telebot.v4"
)

func main() {
	// Initializing .env
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Initializing connection
	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		panic(err)
	}

	//Creating sqlc db
	db := DB.New(conn)

	r := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     os.Getenv("REDIS_USER_PASSWORD"),
		Username:     os.Getenv("REDIS_USER"),
		DB:           0,
		MaxRetries:   5,
		DialTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	})
	if err := r.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	b, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	handlers.SetHandlers(b, db, r, services.CreateUserService(db, r))
	b.Start()
}
