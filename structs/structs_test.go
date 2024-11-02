package structs

import (
	"golearn/utils"
	"testing"
)

func TestRectangle(t *testing.T) {
	t.Run("test the perimeter the Rectangle", func(t *testing.T) {
		a := Rectangle{10.00, 10.00}
		got := a.Perimeter()
		want := 40.0

		if got != want {
			utils.AssertCorrectMessage(t, got, want)
		}
	})

	t.Run("test the area of the Rectangle", func(t *testing.T) {
		a := Rectangle{10.00, 10.00}
		got := a.Area()
		want := 100.00

		if got != want {
			utils.AssertCorrectMessage(t, got, want)
		}
	})
}

func TestCircle(t *testing.T) {
	t.Run("test the perimeter the circle", func(t *testing.T) {
		a := Circle{2.00}
		got := a.Perimeter()
		want := 25.132741228718345

		if got != want {
			utils.AssertCorrectMessage(t, got, want)
		}
	})

	t.Run("test the area of the circle", func(t *testing.T) {
		a := Circle{2.00}
		got := a.Area()
		want := 12.566370614359172

		if got != want {
			utils.AssertCorrectMessage(t, got, want)
		}
	})
}
