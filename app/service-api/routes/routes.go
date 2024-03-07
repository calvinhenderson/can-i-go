package routes

import (
	"io"
	"log"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.Header().Set("cache-control", "no-cache")
		w.WriteHeader(http.StatusOK)

		if _, err := io.WriteString(w, "OK\n"); err != nil {
			log.Printf("Error writing response body")
		}
	}
}
