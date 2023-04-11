package handlinghttprequests

import (
	"fmt"
	"net/http"
	"time"
)

// models
type Product struct {
	Id   int
	Name string `json:"product_name"`
}

var productList []Product = []Product{
	// {1, "product 1"},
	// {2, "product 2"},
	// {3, "product 3"},
	// {4, "product 4"},
	// {5, "product 5"},
}

func ManageHandler() {
	/* genJson() */
	RequestRouteRequest()
	StartListening()
}

// listeners
func StartListening() {
	http.ListenAndServe(":5000", nil)
}

// routers
func RequestRouteRequest() {

	titlesHandler := http.HandlerFunc(getHeaderTitle)
	http.Handle("/Title/", MiddelwareHandler(titlesHandler))

	http.HandleFunc("/Products", productHandler)
	/* http.HandleFunc("/Title/", getHeaderTitle) */
	http.HandleFunc("/Products/", getProductHandler)
	http.Handle("/About", &Product{Name: "this is about page"})
	http.Handle("/", &Product{Name: "this is home page"})
}

// middelware

func MiddelwareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-> before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finishes: %s\n", time.Since(start))
	})

}
