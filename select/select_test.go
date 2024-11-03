package selectme

import (
	"golearn/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	t.Run("test the fast url", func(t *testing.T) {
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			utils.AssertCorrectMessage(t, got, want)
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
