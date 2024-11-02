package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search in a map for known", func(t *testing.T) {
		a := map[string]any{
			"hey":   1,
			"there": 2,
		}

		got, _ := Search(a, "there")
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("search in map for unknown", func(t *testing.T) {
		a := map[string]any{
			"hey":   1,
			"there": 2,
		}

		_, err := Search(a, "know")
		want := "could not find the word you are looking for"

		if err == nil {
			t.Error("expected to get errors")
		}

		if err.Error() != want {
			t.Errorf("err %v, want %v", err, want)
		}
	})
}
