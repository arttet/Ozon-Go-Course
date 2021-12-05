package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var histogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Namespace: "acme",
	Name:      "histogram",
	Help:      "This is my histogram",
	// Buckets: prometheus.DefBuckets, <- само по-умолчанию
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	go func() {
		for {
			histogram.Observe(rand.Float64() * 10)

			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}
