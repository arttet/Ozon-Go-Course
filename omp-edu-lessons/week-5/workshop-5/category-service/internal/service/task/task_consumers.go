package task

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ozonmp/week-5-workshop/category-service/internal/service/database"
)

const (
	consumersCount = 4 * 3
	consumersDelay = 5 * time.Second
)

func (s Service) runTaskConsumers() {
	for i := 0; i < consumersCount; i++ {
		go func() {
			for {
				ctx := context.Background()
				err := s.ExecTask(ctx)
				if err != nil {
					log.Error().Msg("ExecTask()")
				}

				time.Sleep(consumersDelay)
			}
		}()
	}
}

func (s Service) ExecTask(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	var task *Task
	err := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) (txErr error) {
		task, txErr = s.repository.FindNonStartedTask(ctx, tx)
		if txErr != nil {
			return errors.Wrap(txErr, "repository.FindNonStartedTask()")
		}

		// Нету свободных задач
		if task == nil {
			return
		}

		now := time.Now()
		task.StartedAt = &now

		txErr = s.repository.UpdateTask(ctx, task, tx)
		return errors.Wrap(txErr, "repository.UpdateTask()")
	})
	if err != nil {
		return err
	}
	if task == nil {
		return nil
	}

	// Имитация работы
	execDuration := time.Duration(task.ExecDuration)
	time.Sleep(execDuration)

	err = s.repository.DeleteTask(ctx, task, nil)
	return errors.Wrap(err, "repository.DeleteTask()")
}
