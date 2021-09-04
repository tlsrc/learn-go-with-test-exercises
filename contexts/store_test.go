package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (rw *SpyResponseWriter) Header() http.Header {
	return map[string][]string{}
}

func (rw *SpyResponseWriter) Write([]byte) (int, error) {
	rw.written = true
	return 0, nil
}

func (rw *SpyResponseWriter) WriteHeader(statusCode int) {
	rw.written = true
}

// func (s *SpyStore) assertCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("Store should have been cancelled")
// 	}
// }

// func (s *SpyStore) assertNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("Store should not have been cancelled")
// 	}
// }

func TestServer(t *testing.T) {
	data := "response"

	t.Run("Serves data from the store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, req)

		respBody := resp.Body.String()
		if respBody != data {
			t.Errorf(`got "%s", want "%s"`, respBody, data)
		}
	})

	t.Run("Cancels the store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancelFunc := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancelFunc)
		req = req.WithContext(cancellingCtx)

		resp := &SpyResponseWriter{}

		server.ServeHTTP(resp, req)

		if resp.written {
			t.Errorf("No response should have been written due to context cancellation")
		}

	})

}
