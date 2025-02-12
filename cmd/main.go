package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("n")
		if name != "" {
			fmt.Fprintf(w, "Hello, %s", name)
			return
		}
		fmt.Fprint(w, "Hello, gandon!")
	})

	http.ListenAndServe(":12345", r)
}
