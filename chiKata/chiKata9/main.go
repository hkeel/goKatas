package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(jwtauth.Verifier(tokenAuth))

	// Public route
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		w.Write([]byte(tokenString))
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("This is a protected route"))
		})
	})

	http.ListenAndServe(":3000", r)
}

var tokenAuth *jwtauth.JWTAuth

const Secret = "secret"

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(Secret), nil)
}
