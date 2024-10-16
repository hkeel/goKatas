package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/pantry", PantryRoutes())

	http.ListenAndServe(":3000", r)
}

func PantryRoutes() chi.Router {
	r := chi.NewRouter()
	pantryHandler := PantryHandler{}

	r.Get("/", pantryHandler.ListItems)
	r.Post("/", pantryHandler.CreateItem)
	r.Get("/{id}", pantryHandler.GetItem)
	r.Put("/{id}", pantryHandler.UpdateItem)
	r.Delete("/{id}", pantryHandler.DeleteItem)
	return r
}
