package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type Product struct {
// 	Name string `json:"Name"`
// 	Id   int    `json:"Id"`
// }

// var ProductList []Product

// func init() {
// 	productsJSON := `[
// 		{
// 			"Name" : one,
// 			"Id" : 1
// 		},
// 		{
// 			"Name" : two,
// 			"Id" : 2
// 		}
// 	]`
// 	err := json.Unmarshal([]byte(productsJSON), &ProductList)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func productHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		productsJSON, err := json.Marshal(ProductList)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(productsJSON)
// 	}
// }

// func main() {
// 	http.HandleFunc("/products", productHandler)
// 	http.ListenAndServe(":5000", nil)
// }
