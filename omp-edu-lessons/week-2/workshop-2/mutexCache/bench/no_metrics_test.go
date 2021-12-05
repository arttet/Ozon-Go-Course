package bench

import (
	"testing"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/no_metrics/no_mutex"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/no_metrics/with_mutex"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/no_metrics/with_rw_mutex"
)

const parallelFactor = 10_000_00

func Benchmark_NoMutex(b *testing.B) {
	b.Skip("panic in NoMutex")
	c := no_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_Mutex_BalancedLoad(b *testing.B) {
	c := with_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutex_BalancedLoad(b *testing.B) {
	c := with_rw_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}

func Benchmark_Mutex_ReadIntensiveLoad(b *testing.B) {
	c := with_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutex_ReadIntensiveLoad(b *testing.B) {
	c := with_rw_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}
