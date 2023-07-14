package main

import (
	"context"
	"log"
	"os"

	"github.com/ShadrackAdwera/pg-csv/internal/rest"
	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUrl := os.Getenv("DB_URL")

	pool, err := pgxpool.New(context.Background(), dbUrl)

	if err != nil {
		log.Fatal("Error connecting to the DB")
	}

	store := internal.NewStore(pool)

	srv := rest.NewServer(pool, store)

	if err = srv.StartServer(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}

}
