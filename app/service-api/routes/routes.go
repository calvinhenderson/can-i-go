package routes

import (
	"io"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.Header().Set("cache-control", "no-cache")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "OK\n")
	}
}
