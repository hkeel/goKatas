package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/webhook", WebhookHandler)

	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", r)
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received webhook: %+v\n", payload)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received"))
}

type WebhookPayload struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
