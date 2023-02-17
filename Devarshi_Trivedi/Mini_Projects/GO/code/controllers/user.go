package controllers

import (
	"net/http"
	"regexp"
	"training/go/webservice/models"
)

type userController struct {
	userIDP *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, Controller here!"))
}
 
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get (id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u,w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request){
	u, err := uc.parseRequest(r)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc *userController) parseREquest(r *http.Request)

func newUserController() *userController {
	return &userController{
		userIDP: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
