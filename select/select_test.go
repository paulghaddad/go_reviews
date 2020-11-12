package select_go

import (
	"net/http"
  "net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Normal case", func(t *testing.T) {
		slowServer := makeDelayedServer(10)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Did not expect an error: %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected an error")
		}
	})
}


func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
