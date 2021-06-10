package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ch := make(chan string, 1)

		go func() {
			ch <- store.Fetch()
		}()

		select {
		case d := <-ch:
			fmt.Fprint(rw, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
