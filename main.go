package main

import (
	"encoding/json"
	"net/http"
)

// Image ...
type Image struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Product ...
type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
	Image Image   `json:"image"`
}

var products = []Product{Product{0, "book", 10.5, Image{"", 10, 10}}, Product{0, "book", 10.5, Image{"", 10, 10}}}

func main() {
	http.HandleFunc("/product", ShowProduct)
	http.HandleFunc("/", ShowProducts)
	http.ListenAndServe(":8080", nil)
}

// ShowProduct (w http.ResponseWriter, r *http.Request)
func ShowProduct(w http.ResponseWriter, r *http.Request) {
	// products := []Product{}
	// products = append(products, Product{0, "book", 10.5, Image{"", 10, 10}})
	// products = append(products, Product{1, "pencel", 10.5, Image{"", 20, 10}})

	product := Product{0, "book", 10.5, Image{"", 20, 10}}
	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ShowProducts (w http.ResponseWriter, r *http.Request)
func ShowProducts(w http.ResponseWriter, r *http.Request) {
	// products := []Product{Product{0, "book", 10.5, Image{"", 10, 10}}, Product{0, "book", 10.5, Image{"", 10, 10}}}
	// products = append(products, Product{0, "book", 10.5, Image{"", 10, 10}})
	// products = append(products, Product{1, "pencel", 10.5, Image{"", 20, 10}})

	// product := Product{{0, "book", 10.5}, {0, "book1", 10.5}}
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
