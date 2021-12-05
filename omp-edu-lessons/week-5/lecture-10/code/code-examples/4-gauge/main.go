package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "acme",
	Name:      "gauge",
	Help:      "This is my gauge",
})

func init() {
	rand.Seed(time.Now().UnixNano())

	prometheus.MustRegister(gauge)
}

func main() {
	go func() {
		for {
			gauge.Add(rand.Float64()*15 - 5)

			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}
