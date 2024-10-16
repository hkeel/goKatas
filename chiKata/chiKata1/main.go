package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// creates a new router instance
	r := chi.NewRouter()

	// adds a logger middleware to the router
	r.Use(middleware.Logger)

	/* defines a route for the HTTP GET method at the root path "/". When a GET request is made to the root path, the provider handler function is executed.
	This handler function writes the response "Hello, World!" to the client.
	*/
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// starts an HTTP server on port 3000 and uses the chi router to handle requests.
	http.ListenAndServe(":3000", r)
}