package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		if got != want {
			assert.Equal(t, got, want)
		}
	})

	t.Run("when empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		if got != want {
			assert.Equal(t, got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		if got != want {
			assert.Equal(t, got, want)
		}
	})

	t.Run("the langauge is not supported", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Hello, Chris"

		if got != want {
			assert.Equal(t, got, want)
		}
	})
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Chris", "English")
	}
}
