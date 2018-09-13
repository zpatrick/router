package router

import "net/http"

// MethodHandlers map http methods to http.Handlers.
type MethodHandlers map[string]http.Handler

// RouteMap map url path patterns to MethodHandlers.
type RouteMap map[string]MethodHandlers

// ApplyMiddlware applies each middleware to each http.Handler in r.
func (rm RouteMap) ApplyMiddleware(middleware ...Middleware) {
	for path, methodHandlers := range rm {
		for method, _ := range methodHandlers {
			for _, middleware := range middleware {
				rm[path][method] = middleware(rm[path][method])
			}
		}
	}
}

// GlobMatch returns HandlerMatchers for each http.Handler in r using NewGlobHandlerMatcher.
func (rm RouteMap) GlobMatch() []HandlerMatcher {
	return rm.constructMatchers(NewGlobHandlerMatcher)
}

// RegexMatch returns HandlerMatchers for each http.Handler in r using NewRegexHandlerMatcher.
func (rm RouteMap) RegexMatch() []HandlerMatcher {
	return rm.constructMatchers(NewRegexHandlerMatcher)
}

// StringMatch returns HandlerMatchers for each http.Handler in r using NewStringHandlerMatcher.
func (rm RouteMap) StringMatch() []HandlerMatcher {
	return rm.constructMatchers(NewStringHandlerMatcher)
}

// VariableMatch returns HandlerMatchers for each http.Handler in r using NewVariableHandlerMatcher.
func (rm RouteMap) VariableMatch() []HandlerMatcher {
	return rm.constructMatchers(NewVariableHandlerMatcher)
}

func (rm RouteMap) constructMatchers(constructor func(string, string, http.Handler) HandlerMatcher) []HandlerMatcher {
	matchers := []HandlerMatcher{}
	for pattern, methodHandlers := range rm {
		for method, handler := range methodHandlers {
			matchers = append(matchers, constructor(method, pattern, handler))
		}
	}

	return matchers
}
