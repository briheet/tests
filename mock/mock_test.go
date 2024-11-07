package mock

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleep := SpySleeper{}

	CountDown(&buffer, &spySleep)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		assert.Equal(t, got, want)
	}

	if spySleep.Calls != startingNumber {
		t.Errorf("not enough calls to sleeper, want %v, got %v", startingNumber, spySleep.Calls)
	}
}
