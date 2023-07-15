package workers

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5"
	log "github.com/rs/zerolog/log"
)

const TaskSendCsvDataToPg = "task:copy_sales_data_to_pg"

func (d *RedisTaskDistributor) DistroDataOnCsv(ctx context.Context, payload *multipart.FileHeader, options ...asynq.Option) error {
	f, err := payload.Open()

	if err != nil {
		return fmt.Errorf("unable to open file %w", asynq.SkipRetry)
	}

	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return fmt.Errorf("unable to read csv file %w", asynq.SkipRetry)
	}

	data, err := json.Marshal(records)

	if err != nil {
		return fmt.Errorf("unable to marshall data %w", asynq.SkipRetry)
	}

	task := asynq.NewTask(TaskSendCsvDataToPg, data, options...)

	info, err := d.client.EnqueueContext(ctx, task)

	if err != nil {
		return fmt.Errorf("unable to enqueue task context : %w", err)
	}

	log.Info().Str("type", task.Type()).Str("queue_name", info.Queue).Msg("enqueued task")

	return nil
}

func ConvertDateFormat(dateStr string) (string, error) {
	// Parse the input date string using the specified format
	date, err := time.Parse("1/2/2006", dateStr)
	if err != nil {
		return "", err
	}

	// Set the time portion to midnight (00:00:00)
	date = date.Add(time.Duration(-date.Hour()) * time.Hour)
	date = date.Add(time.Duration(-date.Minute()) * time.Minute)
	date = date.Add(time.Duration(-date.Second()) * time.Second)

	// Format the date as a timestamp with time zone (timestamptz)
	formattedTimestamp := date.Format("2006-01-02 15:04:05-07:00")

	return formattedTimestamp, nil
}

func (p *RedisTaskProcessor) ProcessSendCsvDataToDb(
	ctx context.Context,
	task *asynq.Task,
) error {
	var records [][]any

	if err := json.Unmarshal(task.Payload(), &records); err != nil {
		return fmt.Errorf("unable to unmarshall json data %w", asynq.SkipRetry)
	}

	log.Info().Msg("received message from queue . . . ")

	now := time.Now()

	_, err := p.pool.CopyFrom(ctx, pgx.Identifier{"sales"},
		[]string{"region", "country", "item_type", "sales_channel", "order_priority", "order_date", "order_id", "ship_date", "units_sold", "unit_price", "unit_cost", "total_revenue", "total_cost", "total_profit"},
		pgx.CopyFromRows(records))

	if err != nil {
		return fmt.Errorf("")
	}

	log.Info().Any("time to process", time.Since(now)).Msg("all tasks have been processed . . . ")

	return nil
}
