package contexts

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, store.Fetch())
	}
}
