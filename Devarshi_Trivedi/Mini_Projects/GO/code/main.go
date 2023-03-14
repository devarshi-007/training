package main

import (
	"net/http"
	"training/go/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
