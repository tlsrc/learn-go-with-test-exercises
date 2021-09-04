package contexts

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "response"
	server := Server(&StubStore{data})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	respBody := resp.Body.String()
	if respBody != data {
		t.Errorf(`got "%s", want "%s"`, respBody, data)
	}

}
