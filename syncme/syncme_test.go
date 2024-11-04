package syncme

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		wantedCount := 100000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i <= wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.value != want {
		t.Errorf("got %d, want %d", got.value, want)
	}
}
