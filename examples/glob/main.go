package main

import (
	"net/http"

	"github.com/zpatrick/router"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// TODO: maybe this was too extendible?
	// using strict glob isn't good enough - would like :productID
	rm := router.RouteMap{
		"/products": router.MethodHandlers{
			http.MethodGet:  http.HandlerFunc(GetProducts),
			http.MethodPost: http.HandlerFunc(AddProduct),
		},
		"/products/*": router.MethodHandlers{
			http.MethodGet:    http.HandlerFunc(GetProduct),
			http.MethodPut:    http.HandlerFunc(UpdateProduct),
			http.MethodDelete: http.HandlerFunc(DeleteProduct),
		},
	}

	r := router.NewRouter(rm.GlobMatchers())

	// maybe add a new matcher type?
	rm = router.RouteMap{
		"/products": router.MethodHandlers{
			http.MethodGet:  http.HandlerFunc(GetProducts),
			http.MethodPost: http.HandlerFunc(AddProduct),
		},
		"/products/:productID": router.MethodHandlers{
			http.MethodGet:    http.HandlerFunc(GetProduct),
			http.MethodPut:    http.HandlerFunc(UpdateProduct),
			http.MethodDelete: http.HandlerFunc(DeleteProduct),
		},
	}

	r := router.NewRouter(rm.Variable())
	http.ListenAndServe(":9090", r)
}
