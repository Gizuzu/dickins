package database

import (
	"context"
	"log"
	"os"
	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed prepare.sql
var prepare string

var Pool *pgxpool.Pool

func Connect() {
	var err error
	Pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	_, err = Pool.Exec(context.Background(), prepare)
	if err != nil {
		log.Fatal("Failed to prepare database: ", err)
	}
}

func Disconnect() {
	Pool.Close()
}
