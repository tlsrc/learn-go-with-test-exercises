package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		res, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprintf(rw, res)
	}
}
