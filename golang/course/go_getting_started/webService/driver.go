package webService

import "net/http"

func Main() {
	RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
