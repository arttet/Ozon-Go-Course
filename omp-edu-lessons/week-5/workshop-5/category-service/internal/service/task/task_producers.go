package task

import (
	"context"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	producersCount = 4
	producersDelay = 5 * time.Second
)

const (
	jobMinDuration = 5 * time.Second
	jobMaxDuration = 15 * time.Second
)

func (s Service) runTaskProducers() {
	for i := 0; i < producersCount; i++ {
		go func() {
			for {
				ctx := context.Background()
				err := s.CreateNewTask(ctx)
				if err != nil {
					log.Error().Msg("CreateNewTask()")
				}

				time.Sleep(producersDelay)
			}
		}()
	}
}

func (s Service) CreateNewTask(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	minDurationMs := jobMinDuration.Milliseconds()
	maxDurationMs := jobMaxDuration.Milliseconds()
	execDuration := time.Duration(minDurationMs+rand.Int63n(maxDurationMs-minDurationMs)) * time.Millisecond

	task := &Task{
		ExecDuration: Duration(execDuration),
	}

	err := s.repository.InsertTask(ctx, task, nil)
	return errors.Wrap(err, "repository.InsertTask()")
}
