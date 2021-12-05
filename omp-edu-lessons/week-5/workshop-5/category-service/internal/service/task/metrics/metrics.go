package task_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	taskCountTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "category_service",
		Name:      "task_count_total",
		Help:      "Total count of tasks",
	}, []string{"status"})
)

//go:generate stringer -linecomment -type=status
type status uint

const (
	_                = status(iota)
	StatusAll        // all
	StatusNonStarted // non_started
	StatusStarted    // started
)

func SetTaskCountTotal(c uint, s status) {
	taskCountTotal.WithLabelValues(s.String()).Set(float64(c))
}
