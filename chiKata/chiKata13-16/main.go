package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var spotifyClient spotify.Client

type contextKey string

const apiVersionKey contextKey = "apiVersion"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://example.com", "http://localhost:3000"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	initSpotifyClient()

	// Version 2
	r.Route("/v2", func(r chi.Router) {
		r.Use(apiVersionCtx("v2"))
		r.With(validateSpotifyID).Get("/artist/{id}", getArtistHandler)
	})

	// Version 1
	r.Route("/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.With(validateSpotifyID).Get("/artist/{id}", getArtistHandler)
	})

	http.ListenAndServe(":3000", r)
}

func apiVersionCtx(version string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), apiVersionKey, version)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
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

	// small test for api versioning
	version := r.Context().Value(apiVersionKey).(string)
	if version == "v2" {
		artist.Name = strings.ToUpper(artist.Name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}

func validateSpotifyID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		artistID := chi.URLParam(r, "id")
		validID := regexp.MustCompile(`^[a-zA-Z0-9]{22}$`).MatchString(artistID)
		if !validID {
			http.Error(w, "Invalid Spotify ID", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
