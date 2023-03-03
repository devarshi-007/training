package handlinghttprequestwebservices

import (
	"encoding/json"
	"io"
	realdatabase "main/httpRequestWithDatabase/MysqlDatabase"
	"net/http"
	"strconv"
	"strings"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList, err := realdatabase.GetProductList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		productInJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productInJson)
	case http.MethodPost:
		var newProduct realdatabase.Product
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = realdatabase.InsertProduct(newProduct)
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
	product, err := realdatabase.GetProduct(productId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
		realdatabase.RemoveProduct(productId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	case http.MethodPut:
		var updatedProduct realdatabase.Product
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.Id != productId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		realdatabase.UpdatedProduct(updatedProduct)
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
