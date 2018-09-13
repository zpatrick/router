package router

import (
	"net/http"
)

// Router is the root handler for an application.
type Router struct {
	Matchers []HandlerMatcher
	NotFound func(http.ResponseWriter, *http.Request)
}

// NewRouter returns an initialized Router with the specified matchers.
func NewRouter(matchers []HandlerMatcher) *Router {
	return &Router{
		Matchers: matchers,
		NotFound: http.NotFound,
	}
}

// ServeHTTP attempts to match r to a http.Handler using o.Matchers.
// If no match is found, the o.NotFound is executed.
func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, match := range o.Matchers {
		handler, ok := match(r)
		if ok {
			handler.ServeHTTP(w, r)
			return
		}
	}

	o.NotFound(w, r)
}
