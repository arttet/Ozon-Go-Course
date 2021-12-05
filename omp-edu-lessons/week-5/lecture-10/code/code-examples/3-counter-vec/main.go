package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "acme",
	Name:      "counter_vec",
	Help:      "Counter with labels",
}, []string{"label_name"})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for _, labelName := range []string{"a", "b", "c", "d", "e"} {
		go job(labelName)
	}

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}

func job(labelName string) {
	for {
		counter.WithLabelValues(labelName).Inc()

		// Рандомная задержка от 1 до 2 секунд
		delay := time.Duration(1000+rand.Intn(1000)) * time.Millisecond
		time.Sleep(delay)
	}
}
