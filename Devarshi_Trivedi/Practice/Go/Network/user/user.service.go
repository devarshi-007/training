package user

import (
	"encoding/json"
	"fmt"
	"log"
	"main/cors"
	"net/http"
	"strconv"
	"strings"
)

const urlString = "usersdata"

func usersData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userList := getUserList()
		if userList == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(userList)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateUser(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func userData(w http.ResponseWriter, r *http.Request) {
	urlSegment := strings.Split(r.URL.Path, fmt.Sprintf("%s/", urlString))
	if len(urlSegment[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(urlSegment[len(urlSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		user := getUser(userId)
		if user == nil {
			w.Write([]byte("user not found"))
			return
		}
		j, err := json.Marshal(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if user.Id != userId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateUser(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		removeUser(userId)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetUpRoutes(apiBasePath string) {
	UsersHandler := http.HandlerFunc(usersData)
	UserHandler := http.HandlerFunc(userData)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, urlString), cors.Middleware(UsersHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, urlString), cors.Middleware(UserHandler))
}
