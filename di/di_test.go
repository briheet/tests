package di

import (
	"bytes"
	"golearn/utils"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		utils.AssertCorrectMessage(t, got, want)
	}
}
