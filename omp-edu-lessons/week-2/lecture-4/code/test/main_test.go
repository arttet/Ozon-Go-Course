package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_SomeTest функция которая будет запущена во время тестов
func Test_SomeTest(t *testing.T) {
	t.Parallel() // <- запустим тест конкуретно с другими тестами

	t.Run("sum is calculated right", func(t *testing.T) {
		t.Parallel()

		res := CalcSum(1, 2)

		assert.Equal(t, res, res)
	})

	t.Run("error for broken string", func(t *testing.T) {
		t.Parallel()

		res, err := CalcSumFromStrings("text", "3")

		assert.NotNil(t, err)
		t.Log(err)
		assert.Equal(t, 0, res)
	})
}
