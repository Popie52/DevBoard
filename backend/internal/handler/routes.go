package handler

import (
	"fmt"
	"html"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
