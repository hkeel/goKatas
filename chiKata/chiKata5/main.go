package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// 5. Create a route using Chi that accepts POST requests and echoes back the request body

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var body Person
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}