package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "acme",
	Name:      "counter",
	Help:      "This is my counter",
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	go job()

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}

func job() {
	for {
		counter.Inc()

		// Случайная задержка от 1 до 2 секунд
		delay := time.Duration(1000+rand.Intn(1000)) * time.Millisecond
		time.Sleep(delay)
	}
}
