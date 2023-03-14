package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Name string `json:"Name"`
	Id   int    `json:"Id"`
}

var ProductList []Product

func init() {
	productsJSON := `[
		{
			"Name" : "one",
			"Id" : 1
		},
		{
			"Name" : "two",
			"Id" : 2
		}
	]`
	err := json.Unmarshal([]byte(productsJSON), &ProductList)
	if err != nil {
		log.Fatal(err)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product, listItemIdex := findProductById(productId)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)

	case http.MethodPut:
		var updateProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateProduct.Id != productId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		product = &updateProduct
		ProductList[listItemIdex] = *product
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(ProductList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.Id = getNextID()
		ProductList = append(ProductList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return

	}
}

func getNextID() int {
	highestID := -1
	for _, product := range ProductList {
		if highestID < product.Id {
			highestID = product.Id
		}
	}
	return highestID + 1
}

func findProductById(pid int) (*Product, int) {
	for i, product := range ProductList {
		if product.Id == pid {
			return &product, i
		}
	}
	return nil, 0
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middlewear starts")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middlewear finished; %s", time.Since(start))
	})
}

func main() {
	productListenHandler := http.HandlerFunc(productsHandler)
	productItemHandler := http.HandlerFunc(productHandler)
	http.Handle("/products", middlewareHandler(productListenHandler))
	http.Handle("/products/", middlewareHandler(productItemHandler))
	http.ListenAndServe(":5100", nil)
}
