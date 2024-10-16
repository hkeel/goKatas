package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/person/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		person := getPerson(name)
		if person == nil {
			http.Error(w, "Person not found", http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(person)
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

var people = []*Person{
	{"Alice", 25},
	{"Bob", 30},
	{"Charlie", 35},
}

func getPerson(name string) *Person {
	for _, p := range people {
		if strings.EqualFold(p.Name, name) {
			return p
		}
	}
	return nil
}