package handlinghttprequestwebservices

import (
	"fmt"
	realdatabase "main/httpRequestWithDatabase/MysqlDatabase"
	"net/http"
)

const (
	apiBasePath = ""
	apiProducs  = "Products"
)

func ManageHandler() {

	realdatabase.SetupDatabase()

	fmt.Println("database is successfully connected...")

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
