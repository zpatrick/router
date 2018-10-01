package router

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	b := bytes.NewBuffer(nil)
	log.SetOutput(b)

	recorder := httptest.NewRecorder()
	r := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/path"},
	}

	LoggingMiddleware()(handler).ServeHTTP(recorder, r)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.True(t, called)
	assert.True(t, strings.Contains(b.String(), "GET /path"))
}

func TestBasicAuthMiddleware(t *testing.T) {
	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	recorder := httptest.NewRecorder()
	r := &http.Request{Header: http.Header{}}
	r.SetBasicAuth("admin", "pass")

	BasicAuthMiddleware("admin", "pass")(handler).ServeHTTP(recorder, r)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.True(t, called)
}

func TestBasicAuthMiddlewareInvalidAuth(t *testing.T) {
	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	recorder := httptest.NewRecorder()
	r := &http.Request{Header: http.Header{}}
	r.SetBasicAuth("user", "pswrd")

	BasicAuthMiddleware("admin", "pass")(handler).ServeHTTP(recorder, r)
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.False(t, called)
}
