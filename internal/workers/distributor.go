package workers

import (
	"context"
	"mime/multipart"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistroDataOnCsv(ctx context.Context, payload *multipart.FileHeader, options ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(r asynq.RedisConnOpt) TaskDistributor {
	c := asynq.NewClient(r)
	return &RedisTaskDistributor{
		client: c,
	}
}
