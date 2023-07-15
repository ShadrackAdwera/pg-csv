package main

import (
	"context"
	"log"
	"os"

	"github.com/ShadrackAdwera/pg-csv/internal/rest"
	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/ShadrackAdwera/pg-csv/internal/workers"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	zerolog "github.com/rs/zerolog/log"
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

	redisOpts := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_ADDRESS"),
	}

	store := internal.NewStore(pool)
	distro := workers.NewTaskDistributor(redisOpts)
	srv := rest.NewServer(pool, store, distro)

	go runTaskProcessor(redisOpts, pool)

	if err = srv.StartServer(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}

}

func runTaskProcessor(opts asynq.RedisClientOpt, pool *pgxpool.Pool) {
	processor := workers.NewTaskProcessor(opts, pool)

	err := processor.Start()

	if err != nil {
		zerolog.Fatal().Err(err).Msg("unable to start task processor")
	}

	zerolog.Info().Msg("started task processor")
}
