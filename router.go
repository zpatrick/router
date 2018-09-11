package router

import "net/http"

type Router struct {
	NotFound func(http.ResponseWriter, *http.Request)
	// StrictSlash?
	matchers []RouteMatcher
}

func NewRouter(matchers []RouteMatcher) *Router {
	return &Router{
		NotFound: http.NotFound,
		matchers: matchers,
	}
}

func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, match := range o.matchers {
		handler, ok := match(r)
		if ok {
			handler.ServeHTTP(w, r)
			return
		}
	}

	o.NotFound(w, r)
}
