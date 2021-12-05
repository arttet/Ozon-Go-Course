package bench

import (
	"testing"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/with_metrics/with_atomic"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/with_metrics/with_int"
)

func Benchmark_MutexWithMetricsInt(b *testing.B) {
	c := with_int.New()
	for i := 0; i < b.N; i++ {
		emulateLoadWithMetrics(c, parallelFactor)
	}
}

func Benchmark_MutexWithMetricsAtomic(b *testing.B) {
	c := with_atomic.New()
	for i := 0; i < b.N; i++ {
		emulateLoadWithMetrics(c, parallelFactor)
	}
}
