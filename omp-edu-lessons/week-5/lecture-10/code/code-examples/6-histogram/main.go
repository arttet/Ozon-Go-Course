package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var histogramDurationSeconds = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "histogram_duration_seconds", // + _{sum | count | bucket}
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for {
				handlerWithMetrics()

				rps := 7 + rand.Intn(6) // 7 - 13 RPS
				time.Sleep(time.Second / time.Duration(rps))
			}
		}()
	}

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}

func handlerWithMetrics() {
	timeStart := time.Now()
	defer func() {
		elapsedSec := time.Since(timeStart).Seconds()
		histogramDurationSeconds.Observe(elapsedSec)
	}()

	var delay time.Duration
	if rand.Intn(100) >= 96 {
		// Небольшая вероятность, что попадаем сюда
		delay = time.Duration(1500+rand.Intn(2500)) * time.Millisecond // 1.5s - 4s
	} else {
		// Чаще всего используем эту задержку
		delay = time.Duration(100+rand.Intn(400)) * time.Millisecond // 100ms - 500ms
	}

	time.Sleep(delay)
}
