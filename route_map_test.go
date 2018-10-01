package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteMapIterate(t *testing.T) {
	rm := RouteMap{
		"/products": MethodHandlers{
			http.MethodGet:  nil,
			http.MethodPost: nil,
		},
		"/products/:productID": MethodHandlers{
			http.MethodGet:    nil,
			http.MethodPut:    nil,
			http.MethodDelete: nil,
		},
	}

	rm.Iterate(func(pattern, method string, handler http.Handler) {
		delete(rm[pattern], method)
		if len(rm[pattern]) == 0 {
			delete(rm, pattern)
		}
	})

	assert.Len(t, rm, 0)
}

func TestRouteMapApplyMiddlware(t *testing.T) {
	rm := RouteMap{
		"/products": MethodHandlers{
			http.MethodGet:  nil,
			http.MethodPost: nil,
		},
		"/products/:productID": MethodHandlers{
			http.MethodGet:    nil,
			http.MethodPut:    nil,
			http.MethodDelete: nil,
		},
	}

	var calls int
	middleware := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
		})
	}

	rm.ApplyMiddleware(middleware)
	rm.Iterate(func(pattern, method string, handler http.Handler) {
		handler.ServeHTTP(nil, nil)
	})

	assert.Equal(t, 5, calls)
}
