package router

import (
	"net/http"
	"regexp"
	"strings"

	glob "github.com/ryanuber/go-glob"
)

type RouteMatcher func(r *http.Request) (http.Handler, bool)

func NewGlobRouteMatcher(method, pattern string, handler http.Handler) RouteMatcher {
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && glob.Glob(pattern, r.URL.Path) {
			return handler, true
		}

		return nil, false
	}
}

func NewRegexRouteMatcher(method, pattern string, handler http.Handler) RouteMatcher {
	re := regexp.MustCompile(pattern)
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && re.MatchString(r.URL.Path) {
			return handler, true
		}

		return nil, false
	}
}

func NewStringRouteMatcher(method, pattern string, handler http.Handler) RouteMatcher {
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && r.URL.Path == pattern {
			return handler, true
		}

		return nil, false
	}
}

func NewVariableRouteMatcher(method, pattern string, handler http.Handler) RouteMatcher {
	patternSegments := Segments(pattern)
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method != method {
			return nil, false
		}

		pathSegments := Segments(r.URL.Path)
		if len(pathSegments) != len(patternSegments) {
			return nil, false
		}

		for i := 0; i < len(patternSegments); i++ {
			if strings.HasPrefix(patternSegments[i], ":") {
				continue
			}

			if pathSegments[i] != patternSegments[i] {
				return nil, false
			}
		}

		return handler, true
	}
}
