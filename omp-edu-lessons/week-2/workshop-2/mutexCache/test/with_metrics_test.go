package test

import (
	"testing"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/impl/with_metrics/with_atomic"

	"github.com/stretchr/testify/assert"
)

func Test_CacheWithMetrics(t *testing.T) {
	t.Parallel()
	// Разные имплементации кешей
	testCache := with_atomic.New()
	//testCache := no_mutex.New()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storedValue)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoadWithMetrics(t, testCache, parallelFactor)
	})
}
