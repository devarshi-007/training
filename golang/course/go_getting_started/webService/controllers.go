package webService

import (
	"net/http"
	"regexp"
)

// defineing class attributes
type userController struct {
	userIDPattern *regexp.Regexp
}

// defining constructor return new object of pointer
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

// defining methods
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bytes comming from server..."))
}
