package handlinghttprequestwebservices

import (
	"encoding/json"
	"fmt"
	"io"
	database "main/handlingHttpRequest.webservices/Database"
	"net/http"
	"strconv"
	"strings"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList := database.GetProductList()
		productInJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productInJson)
	case http.MethodPost:
		var newProduct database.Product
		bodyBytes, err := io.ReadAll(r.Body)
		fmt.Println(string(bodyBytes))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		fmt.Println(newProduct, err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = database.AddOrUpdateProduct(newProduct)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		return
	}
}

// getProductHAndler for /Products/

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegs := strings.SplitAfter(r.URL.Path, "/Products/")
	productId, err := strconv.Atoi(urlPathSegs[len(urlPathSegs)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product := database.GetProduct(productId)
	if product == nil {
		if r.Method == http.MethodDelete {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}
	switch r.Method {
	case http.MethodGet:
		productInJson, err := json.Marshal(*product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productInJson)

	case http.MethodDelete:
		database.RemoveProduct(productId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	case http.MethodPut:
		var updatedProduct database.Product
		bodyBytes, err := io.ReadAll(r.Body)
		fmt.Println(string(bodyBytes))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)
		fmt.Println(updatedProduct.Id, productId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.Id != productId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		database.AddOrUpdateProduct(updatedProduct)
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
