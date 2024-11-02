package utils

import "testing"

func AssertCorrectMessage(t *testing.T, got, want interface{}) {
	t.Helper()

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
