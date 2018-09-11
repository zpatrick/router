package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegments(t *testing.T) {
	cases := map[string][]string{
		"/":             []string{""},
		"/products":     []string{"products"},
		"/products/":    []string{"products", ""},
		"/products/p1":  []string{"products", "p1"},
		"/products/p1/": []string{"products", "p1", ""},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, Segments(input))
		})
	}
}
