package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func isPowerOfTwo(n int) bool {
	a := 1
	for i := 0; i <= n; i++ {
		if a == n {
			return true
		} else {
			a *= 2
		}
	}
	return false
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 128})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	addr := ":3334"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v\n", claims["user_id"])))
			t := int(claims["user_id"].(float64))
			fmt.Println(t)

			if isPowerOfTwo(t) {
				w.Write([]byte(fmt.Sprintf("User_id is power of two")))
			} else {
				w.Write([]byte(fmt.Sprintf("User_id is not power of two")))
			}
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
		})
	})

	return r
}
