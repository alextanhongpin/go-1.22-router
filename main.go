package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// https://pkg.go.dev/net/http#hdr-Patterns

	// This matches only the path / and nothing else.
	// Otherwise, it will match all paths.
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello, world! %s", id)
	})
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "creating users")
	})
	http.ListenAndServe(":8080", mux)
}
