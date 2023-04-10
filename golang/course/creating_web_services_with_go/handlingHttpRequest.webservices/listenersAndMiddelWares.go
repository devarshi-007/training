package handlinghttprequestwebservices

import (
	"fmt"
	database "main/handlingHttpRequest.webservices/Database"
	"net/http"
)

const (
	apiBasePath = ""
	apiProducs  = "Products"
)

func ManageHandler() {
	/* genJson() */
	database.Init()
	RequestRouteRequest()
	StartListening()
}

// listeners
func StartListening() {
	http.ListenAndServe(":5000", nil)
}

// routers
func RequestRouteRequest() {
	handleProducts := http.HandlerFunc(productHandler)
	handleProduct := http.HandlerFunc(getProductHandler)
	http.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, apiProducs), handleProducts)
	http.HandleFunc(fmt.Sprintf("%s/%s/", apiBasePath, apiProducs), handleProduct)
}
