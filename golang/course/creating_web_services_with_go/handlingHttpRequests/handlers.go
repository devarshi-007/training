package handlinghttprequests

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// handlers functions

// serverHttp for ~/About and ~/
func (f *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Name))
}

// productHandler for /Products
func newtId() int {
	lent := len(productList)
	if lent == 0 {
		return 1
	}
	lastOne := productList[lent-1].Id
	return lastOne + 1
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productInJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productInJson)
	case http.MethodPost:
		var newProduct Product
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
		if newProduct.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.Id = newtId()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// getProductHAndler for /Products/
func getProductById(id int) (int, *Product) {
	for ind, prd := range productList {
		if prd.Id == id {
			return ind, &prd
		}
	}
	return -1, nil
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegs := strings.SplitAfter(r.URL.Path, "/Products/")
	productId, err := strconv.Atoi(urlPathSegs[len(urlPathSegs)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	productListIndex, product := getProductById(productId)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		productInJson, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productInJson)
	case http.MethodDelete:
		if len(productList)-1 == productListIndex {
			productList = productList[:productListIndex]
		} else {
			productList = append(productList[:productListIndex], productList[productListIndex+1:]...)
		}
		w.WriteHeader(http.StatusAccepted)
	case http.MethodPut:
		var updatedProduct Product
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
		product = &updatedProduct
		productList[productListIndex] = *product
		w.WriteHeader(http.StatusCreated)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// getHeaderTitle for /Title/
func getHeaderTitle(w http.ResponseWriter, r *http.Request) {

	urlPathSegs := strings.SplitAfter(r.URL.Path, "/Title/")
	headerSize, err := strconv.Atoi(urlPathSegs[len(urlPathSegs)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	msg := "this is something which you can in help page"

	switch headerSize {
	case 0:
		msg = "<b>" + msg + "</b>"
	case 1:
		msg = "<h1>" + msg + "</h1>"
	case 2:
		msg = "<h2>" + msg + "</h2>"
	case 3:
		msg = "<h3>" + msg + "</h3>"
	case 4:
		msg = "<h4>" + msg + "</h4>"
	case 5:
		msg = "<h5>" + msg + "</h5>"
	case 6:
		msg = "<h6>" + msg + "</h6>"
	default:
		msg = "<p>" + msg + "</p>"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(msg))
}
