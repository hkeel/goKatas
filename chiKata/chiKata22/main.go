package main

// Note: use postman collection to hit /data endpoint then check the response headers. Content-Encoding: gzip should be present.

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Use the compressor middleware to enable response compression
	r.Use(middleware.Compress(5)) // Compression level 5 (range: 1-9)

	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := "This is some data that will be compressed if the client supports it."
		log.Println("Serving /data endpoint")
		w.Write([]byte(data))
	})

	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", r)
}
