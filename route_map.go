package router

import "net/http"

// MethodHandlers map http methods to http.Handlers.
type MethodHandlers map[string]http.Handler

// A RouteMap maps url path patterns to MethodHandlers.
type RouteMap map[string]MethodHandlers

// ApplyMiddlware applies each middleware to each http.Handler in rm.
func (rm RouteMap) ApplyMiddleware(middleware ...Middleware) {
	for path, methodHandlers := range rm {
		for method, _ := range methodHandlers {
			for _, middleware := range middleware {
				rm[path][method] = middleware(rm[path][method])
			}
		}
	}
}

// GlobMatch return a HandlerMatcher for each http.Handler in rm using NewGlobHandlerMatcher.
func (rm RouteMap) GlobMatch() []HandlerMatcher {
	matchers := []HandlerMatcher{}
	rm.Iterate(func(pattern, method string, handler http.Handler) {
		matchers = append(matchers, NewGlobHandlerMatcher(method, pattern, handler))
	})

	return matchers
}

// RegexMatch returns a HandlerMatcher for each http.Handler in rm using NewRegexHandlerMatcher.
func (rm RouteMap) RegexMatch() []HandlerMatcher {
	matchers := []HandlerMatcher{}
	rm.Iterate(func(pattern, method string, handler http.Handler) {
		matchers = append(matchers, NewRegexHandlerMatcher(method, pattern, handler))
	})

	return matchers
}

// StringMatch return a HandlerMatcher for each http.Handler in rm using NewStringHandlerMatcher.
func (rm RouteMap) StringMatch() []HandlerMatcher {
	matchers := []HandlerMatcher{}
	rm.Iterate(func(pattern, method string, handler http.Handler) {
		matchers = append(matchers, NewStringHandlerMatcher(method, pattern, handler))
	})

	return matchers
}

// VariableMatch returns a HandlerMatcher for each http.Handler in rm using NewVariableHandlerMatcher.
func (rm RouteMap) VariableMatch() []HandlerMatcher {
	matchers := []HandlerMatcher{}
	rm.Iterate(func(pattern, method string, handler http.Handler) {
		matchers = append(matchers, NewVariableHandlerMatcher(method, pattern, handler))
	})

	return matchers
}

// Iterate calls fn for each handler in rm.
func (rm RouteMap) Iterate(fn func(pattern, method string, handler http.Handler)) {
	for pattern, methodHandlers := range rm {
		for method, handler := range methodHandlers {
			fn(pattern, method, handler)
		}
	}
}
