package mock

import (
	"fmt"
	"io"
)

const (
	finalword      = "Go!"
	startingNumber = 3
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDown(out io.Writer, spy *SpySleeper) {
	for i := startingNumber; i > 0; i-- {
		fmt.Fprintln(out, i)
		spy.Sleep()
	}
	fmt.Fprint(out, finalword)
}
