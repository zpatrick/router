package router

import (
	"net/http"
	"regexp"
	"strings"

	glob "github.com/ryanuber/go-glob"
)

// A HandlerMatcher is a function that matches a *http.Request to a http.Handler.
type HandlerMatcher func(r *http.Request) (hander http.Handler, matchFound bool)

// NewGlobHandlerMatcher returns a HandlerMatcher that returns a match if and only if
// the request.Method matches method,
// and the request.URL.Path glob matches pattern.
func NewGlobHandlerMatcher(method, pattern string, handler http.Handler) HandlerMatcher {
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && glob.Glob(pattern, r.URL.Path) {
			return handler, true
		}

		return nil, false
	}
}

// NewRegexHandlerMatcher returns a HandlerMatcher that returns a match if and only if
// the request.Method matches method,
// and the request.URL.Path regex matches pattern.
func NewRegexHandlerMatcher(method, pattern string, handler http.Handler) HandlerMatcher {
	re := regexp.MustCompile(pattern)
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && re.MatchString(r.URL.Path) {
			return handler, true
		}

		return nil, false
	}
}

// NewStringHandlerMatcher returns a HandlerMatcher that returns a match if and only if
// the request.Method matches method,
// and the request.URL.Path matches pattern.
func NewStringHandlerMatcher(method, pattern string, handler http.Handler) HandlerMatcher {
	return func(r *http.Request) (http.Handler, bool) {
		if r.Method == method && r.URL.Path == pattern {
			return handler, true
		}

		return nil, false
	}
}

// NewVariableHandlerMatcher returns a HandlerMatcher that returns a match if and only if
// the request.Method matches method,
// and the request.URL.Path variable matches pattern.
// Path variables are specified in pattern by placing a ':' in front of the variable name.
// This is just for human-readability, it denotes that any value can be used
// in the specified segment. Path variables can be fetched using the Segment helper functions.
// Note that the following are functionally equivalent:
//   NewVariableHandlerMatcher(http.MethodGet, "/product/:productID/", handler)
//   NewGlobHandlerMatcher(http.MethodGet, "/product/*/", handler)
func NewVariableHandlerMatcher(method, pattern string, handler http.Handler) HandlerMatcher {
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
