package task

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	task_metrics "github.com/ozonmp/week-5-workshop/category-service/internal/service/task/metrics"
)

type Metrics struct {
	AllCount        uint `db:"all_count"`
	NonStartedCount uint `db:"non_started_count"`
	StartedCount    uint `db:"started_count"`
}

func (s Service) runTaskMetrics() {
	go func() {
		for {
			ctx := context.Background()
			metrics, err := s.repository.GetMetrics(ctx)
			if err != nil {
				log.Error().Msg("repository.GetMetrics()")
			} else {
				task_metrics.SetTaskCountTotal(metrics.AllCount, task_metrics.StatusAll)
				task_metrics.SetTaskCountTotal(metrics.NonStartedCount, task_metrics.StatusNonStarted)
				task_metrics.SetTaskCountTotal(metrics.StartedCount, task_metrics.StatusStarted)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
