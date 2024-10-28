package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// 3. Create a route that accepts query parameters using Chi and returns them in the response
// 4. Implement Chi middleware to log incoming HTTP requests

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello, " + name))
	})

	http.ListenAndServe(":3000", r)
}