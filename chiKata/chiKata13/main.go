package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var spotifyClient spotify.Client

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	initSpotifyClient()

	// Route to get artist information
	r.Get("/artist/{id}", getArtistHandler)

	http.ListenAndServe(":3000", r)
}

func initSpotifyClient() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	httpClient := config.Client(context.Background())
	spotifyClient = spotify.NewClient(httpClient)
}

func getArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")

	artist, err := spotifyClient.GetArtist(spotify.ID(artistID))
	if err != nil {
		http.Error(w, "Error fetching artist information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}