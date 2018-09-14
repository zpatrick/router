package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/zpatrick/router"
)

type Product struct{}

var Products = map[string]Product{
	"p1": Product{},
	"p2": Product{},
	"p3": Product{},
}

func newProductID() string {
	return strconv.Itoa(rand.Int())
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	Products[newProductID()] = product
	w.WriteHeader(200)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := router.Segment(r.URL.Path, 1)
	product, ok := Products[productID]
	if !ok {
		msg := fmt.Sprintf("Product %s does not exist", productID)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := router.Segment(r.URL.Path, 1)
	if _, ok := Products[productID]; !ok {
		msg := fmt.Sprintf("Product %s does not exist", productID)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	delete(Products, productID)
	w.WriteHeader(200)
}

func main() {
	rm := router.RouteMap{
		"/products": router.MethodHandlers{
			http.MethodGet:  http.HandlerFunc(ListProducts),
			http.MethodPost: http.HandlerFunc(AddProduct),
		},
		"/products/:productID": router.MethodHandlers{
			http.MethodGet:    http.HandlerFunc(GetProduct),
			http.MethodDelete: http.HandlerFunc(DeleteProduct),
		},
	}

	rm.ApplyMiddleware(router.LoggingMiddleware())
	r := router.NewRouter(rm.VariableMatch())
	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
