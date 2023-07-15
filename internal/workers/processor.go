package workers

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/rs/zerolog/log"
)

type TaskProcessor interface {
	Start() error
	ProcessSendCsvDataToDb(
		ctx context.Context,
		task *asynq.Task,
	) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	pool   *pgxpool.Pool
}

func NewTaskProcessor(opt asynq.RedisConnOpt, pool *pgxpool.Pool) TaskProcessor {
	server := asynq.NewServer(opt, asynq.Config{
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Str("task_type", task.Type()).Bytes("payload", task.Payload()).Msg("task processing failed")
		}),
	})
	return &RedisTaskProcessor{
		server, pool,
	}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendCsvDataToPg, p.ProcessSendCsvDataToDb)

	return p.server.Start(mux)
}
