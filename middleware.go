package router

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func BasicAuthMiddleware(username, password string) Middleware {
	hash := func(user, pass string) string {
		sum := sha256.Sum256([]byte(user + pass))
		return fmt.Sprintf("%x", sum)
	}

	key := hash(username, password)
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if ok && hash(user, pass) == key {
				handler.ServeHTTP(w, r)
			}

			w.Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte("401 Unauthorized\n")); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
}
