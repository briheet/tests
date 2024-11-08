package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating the character", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"

		if got != want {
			assert.Equal(t, got, want)
		}
	})

	t.Run("no character", func(t *testing.T) {
		got := Repeat("")
		want := ""

		if got != want {
			assert.Equal(t, got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
