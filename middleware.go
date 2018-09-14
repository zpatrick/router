package router

import (
	"crypto/sha256"
	"log"
	"net/http"
)

// Middleware is a function that adds functionality to a handler.
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware returns a Middleware that logs requests' methods and paths.
func LoggingMiddleware() Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	}
}

// BasicAuthMiddleware returns a Middleware that requires the specified
// username password combination to be used in requests' basic auth headers
// before the original handler is executed.
// Otherwise, a 401 Status Unauthorized response is returned.
func BasicAuthMiddleware(username, password string) Middleware {
	sum := sha256.Sum256([]byte(username + password))
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if ok && sha256.Sum256([]byte(user+pass)) == sum {
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
