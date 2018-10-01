package router

import (
	"net/http"
	"net/url"
	"testing"
)

func NewRequest(method, path string) *http.Request {
	return &http.Request{
		Method: method,
		URL:    &url.URL{Path: path},
	}
}

func TestGlobHandlerMatcher(t *testing.T) {
	cases := map[string]struct {
		Matcher  HandlerMatcher
		Request  *http.Request
		Expected bool
	}{
		"Star-only Match": {
			Matcher:  NewGlobHandlerMatcher("GET", "*", nil),
			Request:  NewRequest("GET", "/"),
			Expected: true,
		},
		"Star-only Mismatch (method)": {
			Matcher:  NewGlobHandlerMatcher("GET", "*", nil),
			Request:  NewRequest("PUT", "/"),
			Expected: false,
		},
		"Star-start Match": {
			Matcher:  NewGlobHandlerMatcher("GET", "*/products", nil),
			Request:  NewRequest("GET", "/api/v1/products"),
			Expected: true,
		},
		"Star-start Mismatch (spelling)": {
			Matcher:  NewGlobHandlerMatcher("GET", "*/products", nil),
			Request:  NewRequest("GET", "/api/v1/product"),
			Expected: false,
		},
		"Star-start Mismatch (too long)": {
			Matcher:  NewGlobHandlerMatcher("GET", "*/products", nil),
			Request:  NewRequest("GET", "/api/v1/products/p123"),
			Expected: false,
		},
		"Star-end Match": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products/*", nil),
			Request:  NewRequest("GET", "/products/p123"),
			Expected: true,
		},
		"Star-end Mismatch (too short)": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products/*", nil),
			Request:  NewRequest("GET", "/products"),
			Expected: false,
		},
		"Star-mid Match": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products/*/count", nil),
			Request:  NewRequest("GET", "/products/p123/count"),
			Expected: true,
		},
		"Star-mid Mismatch (spelling)": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products/*/count", nil),
			Request:  NewRequest("GET", "/products/p123/counts"),
			Expected: false,
		},
		"Static Match": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products", nil),
			Request:  NewRequest("GET", "/products"),
			Expected: true,
		},
		"Static Mismatch (spelling)": {
			Matcher:  NewGlobHandlerMatcher("GET", "/products", nil),
			Request:  NewRequest("GET", "/product"),
			Expected: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if _, result := c.Matcher(c.Request); result != c.Expected {
				t.Errorf("Result was %v, expected %v", result, c.Expected)
			}
		})
	}
}

func TestRegexHandlerMatcher(t *testing.T) {

}

func TestStringHandlerMatcher(t *testing.T) {

}

func TestVariableHandlerMatcher(t *testing.T) {

}
