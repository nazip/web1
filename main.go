package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

var products = []Product{
	Product{0, "book", 10.5, Image{"", 10, 10}},
	Product{1, "book1", 10.5, Image{"", 10, 10}},
	Product{2, "book2", 10.5, Image{"", 10, 10}},
	Product{3, "book3", 10.5, Image{"", 10, 10}},
	Product{4, "book4", 10.5, Image{"", 10, 10}},
	Product{5, "book5", 10.5, Image{"", 10, 10}},
	Product{6, "book6", 10.5, Image{"", 10, 10}}}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := httprouter.New()
	router.GET("/products", ShowProducts)
	router.GET("/product/:id", ShowProduct)
	router.GET("/error", ShowError)
	http.ListenAndServe(":"+port, router)
}

// ShowError (w http.ResponseWriter, r *http.Request)
func ShowError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	js, err := json.Marshal("NOT FOUND")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ShowProduct (w http.ResponseWriter, r *http.Request)
func ShowProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, error := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if error != nil {
		http.Redirect(w, r, "", http.StatusNotFound)
		return
	}

	product := products[index]

	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ShowProducts (w http.ResponseWriter, r *http.Request)
func ShowProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
