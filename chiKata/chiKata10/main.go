package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/books", listBooksHandler)

	http.ListenAndServe(":3000", r)
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []*Book{
	{
		ID:     "1",
		Title:  "Safe Haven",
		Author: "Nicholas Sparks",
	},
	{
		ID:     "2",
		Title:  "The Notebook",
		Author: "Nicholas Sparks",
	},
	{
		ID:     "3",
		Title:  "The Last Song",
		Author: "Nicholas Sparks",
	},
	{
		ID:     "4",
		Title:  "Charlotte's Web",
		Author: "E.B. White",
	},
	{
		ID:     "5",
		Title:  "Stuart Little",
		Author: "E.B. White",
	},
	{
		ID:     "6",
		Title:  "The Cat in the Hat",
		Author: "Dr. Seuss",
	},
	{
		ID:     "7",
		Title:  "Green Eggs and Ham",
		Author: "Dr. Seuss",
	},
}

func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 5
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(books) {
		start = len(books)
	}
	if end > len(books) {
		end = len(books)
	}

	paginatedBooks := books[start:end]

	json.NewEncoder(w).Encode(paginatedBooks)
}
