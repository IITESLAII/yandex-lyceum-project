package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	UserName string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"123"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	DbName   string `env:"POSTGRES_DB" env-default:"yandex"`
}

type DB struct {
	Db *sqlx.DB
}

func New(config Config) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.UserName, config.Password, config.DbName, config.Host, config.Port)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Conn(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %w", err)
	}
	return &DB{Db: db}, nil
}
