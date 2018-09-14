package router

import (
	"fmt"
	"net/http"
	"net/url"
)

func ExampleSegments() {
	r := &http.Request{
		URL: &url.URL{Path: "/products/p582"},
	}

	fmt.Println(Segments(r.URL.Path))
	// Output: [products p582]
}

func ExampleSegment() {
	r := &http.Request{
		URL: &url.URL{Path: "/products/p582"},
	}

	fmt.Println(Segment(r.URL.Path, 1))
	// Output: p582
}

func ExampleIntSegment() {
	r := &http.Request{
		URL: &url.URL{Path: "/products/582"},
	}

	productID, _ := IntSegment(r.URL.Path, 1)
	fmt.Println(productID)
	// Output: 582
}

func ExampleNewGlobHandlerMatcher() {
	matcher := NewGlobHandlerMatcher(http.MethodGet, "/products/*/", nil)
	r := &http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{Path: "/products/p582/"},
	}

	if _, ok := matcher(r); ok {
		fmt.Println("Match successful!")
	}

	// Output: Match successful!
}

func ExampleNewRegexHandlerMatcher() {
	matcher := NewRegexHandlerMatcher(http.MethodGet, "/products/.+/", nil)
	r := &http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{Path: "/products/p582/"},
	}

	if _, ok := matcher(r); ok {
		fmt.Println("Match successful!")
	}

	// Output: Match successful!
}

func ExampleNewStringHandlerMatcher() {
	matcher := NewStringHandlerMatcher(http.MethodGet, "/home", nil)
	r := &http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{Path: "/home"},
	}

	if _, ok := matcher(r); ok {
		fmt.Println("Match successful!")
	}

	// Output: Match successful!
}

func ExampleNewVariableHandlerMatcher() {
	matcher := NewVariableHandlerMatcher(http.MethodGet, "/products/:productID", nil)
	r := &http.Request{
		Method: http.MethodGet,
		URL:    &url.URL{Path: "/products/p582"},
	}

	if _, ok := matcher(r); ok {
		fmt.Println("Match successful!")
	}

	// Output: Match successful!
}

func ExampleMiddleware() {
	myMiddleware := func(h http.Handler) http.Handler {
		// do some logic here
		return h
	}

	rm := RouteMap{}
	rm.ApplyMiddleware(myMiddleware)
}

func ExampleLoggingMiddleware() {
	rm := RouteMap{}
	rm.ApplyMiddleware(LoggingMiddleware())
}

func ExampleBasicAuthMiddleware() {
	rm := RouteMap{}
	rm.ApplyMiddleware(BasicAuthMiddleware("admin", "password"))
}

func ExampleRouteMap_GlobMatch() {
	rm := RouteMap{
		"/products": MethodHandlers{
			http.MethodGet:  http.HandlerFunc(nil),
			http.MethodPost: http.HandlerFunc(nil),
		},
		"/products/*/": MethodHandlers{
			http.MethodGet:    http.HandlerFunc(nil),
			http.MethodDelete: http.HandlerFunc(nil),
		},
	}

	r := NewRouter(rm.GlobMatch())
	http.Handle("/", r)
}

func ExampleRouteMap_RegexMatch() {
	rm := RouteMap{
		"/products": MethodHandlers{
			http.MethodGet:  http.HandlerFunc(nil),
			http.MethodPost: http.HandlerFunc(nil),
		},
		"/products/.+/": MethodHandlers{
			http.MethodGet:    http.HandlerFunc(nil),
			http.MethodDelete: http.HandlerFunc(nil),
		},
	}

	r := NewRouter(rm.RegexMatch())
	http.Handle("/", r)
}

func ExampleRouteMap_StringMatch() {
	rm := RouteMap{
		"/home": MethodHandlers{
			http.MethodGet: http.HandlerFunc(nil),
		},
		"/account": MethodHandlers{
			http.MethodGet:  http.HandlerFunc(nil),
			http.MethodPost: http.HandlerFunc(nil),
			http.MethodPut:  http.HandlerFunc(nil),
		},
	}

	r := NewRouter(rm.StringMatch())
	http.Handle("/", r)
}

func ExampleRouteMap_VariableMatch() {
	rm := RouteMap{
		"/products": MethodHandlers{
			http.MethodGet:  http.HandlerFunc(nil),
			http.MethodPost: http.HandlerFunc(nil),
		},
		"/products/:productID": MethodHandlers{
			http.MethodGet:    http.HandlerFunc(nil),
			http.MethodDelete: http.HandlerFunc(nil),
		},
	}

	r := NewRouter(rm.VariableMatch())
	http.Handle("/", r)
}

func ExampleRouter() {
	rm := RouteMap{}
	r := NewRouter(rm.StringMatch())
	http.Handle("/", r)
}
