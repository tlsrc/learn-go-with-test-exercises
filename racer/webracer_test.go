package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("race 2 servers", func(t *testing.T) {
		slowServer := makeHTTPServer(20 * time.Millisecond)
		fastServer := makeHTTPServer(1 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Did not expect an error but got %v", err)
		}
		if got != want {
			t.Errorf("Wanted %v but got %v", want, got)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		timeoutServer := makeHTTPServer(50 * time.Millisecond)

		defer timeoutServer.Close()

		_, err := ConfigurableRacer(timeoutServer.URL, timeoutServer.URL, 25*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didnot get one")
		}
	})

}

func makeHTTPServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}
