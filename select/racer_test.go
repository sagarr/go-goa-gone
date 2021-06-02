package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares the speed of servers, returns the url of fastest one", func(t *testing.T) {
		slowServer := createDelayedServer(10 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatal("didn't expect error")
		}

		if got != fastServer.URL {
			t.Errorf("got %v, want %v", got, fastServer.URL)
		}
	})

	t.Run("return an error if server doesn't respond within the specified time", func(t *testing.T) {
		server := createDelayedServer(11 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error")
		}
	})

}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
