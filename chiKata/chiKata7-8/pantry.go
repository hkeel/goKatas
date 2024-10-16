package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// 8. Implement error handling in a Chi API to return appropriate error responses

type PantryHandler struct{}

func (ph PantryHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listFoodItems())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (ph PantryHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	foodItem := getFoodItem(id)
	if foodItem == nil {
		http.Error(w, "Food Item not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(foodItem)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (ph PantryHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var foodItem FoodItem
	err := json.NewDecoder(r.Body).Decode(&foodItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	foodItem = storeFoodItem(foodItem)
	err = json.NewEncoder(w).Encode(foodItem)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (ph PantryHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var foodItem FoodItem
	err = json.NewDecoder(r.Body).Decode(&foodItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedFoodItem := updateFoodItem(id, foodItem)
	if updatedFoodItem == nil {
		http.Error(w, "Food Item not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedFoodItem)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (ph PantryHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	deletedFoodItem := deleteFoodItem(id)
	if deletedFoodItem == nil {
		http.Error(w, "Food Item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
