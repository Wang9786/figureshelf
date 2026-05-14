package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/redis/go-redis/v9"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgres(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewRedis(redisAddr string, redisURL string) *redis.Client {
	if redisURL != "" {
		options, err := redis.ParseURL(redisURL)
		if err == nil {
			return redis.NewClient(options)
		}
	}

	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}