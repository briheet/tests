package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			assert.Equal(t, got, want)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		got := Sum(a)
		want := 15

		if got != want {
			assert.Equal(t, got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all the arrays given", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 3}, []int{3, 5})
		want := 16

		if got != want {
			assert.Equal(t, got, want)
		}
	})
}

func BenchmarkInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 2)
	}
}
