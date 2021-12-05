package bench

import "testing"

func Benchmark_MutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MutexCounter()
	}
}

func Benchmark_AtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}
