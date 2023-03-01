package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int
	Name      string
	Age       int
	BirthDate string
	Male      string
}

func ShowUser(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/userDb")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	get, err := db.Query(`SELECT Id,Name,Age,BirthDate,Male FROM user`)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	userList := make([]User, 0)

	for get.Next() {
		var user User
		get.Scan(&user.Id, &user.Name, &user.Age, &user.BirthDate, &user.Male)
		userList = append(userList, user)
	}
	data, err := json.Marshal(userList)
	defer get.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Write(data)

	defer db.Close()

}

func ShowUserWithId(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/userDb")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	defer db.Close()

	switch r.Method {
	case http.MethodGet:
		u := strings.Split(r.URL.Path, "user/")
		Id, err := strconv.Atoi(u[len(u)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		get := db.QueryRow(`SELECT Id,Name,Age,BirthDate,Male FROM user WHERE Id = ?`, Id)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		var user User
		get.Scan(&user.Id, &user.Name, &user.Age, &user.BirthDate, &user.Male)
		data, err := json.Marshal(user)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write(data)

	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)
		if err != nil {
			fmt.Println("1")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result, err := db.Exec(`INSERT into user
		(Name,Age,BirthDate) VALUES (?,?,?)`, user.Name, user.Age, user.BirthDate)

		insid, _ := result.LastInsertId()
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := json.Marshal(insid)
		if err != nil {
			fmt.Println("3")
			fmt.Println(err)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		w.Write(id)

	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		u := strings.Split(r.URL.Path, "/user/")
		Id, err := strconv.Atoi(u[len(u)-1])
		fmt.Println(Id)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, err = db.Exec(`UPDATE user SET Name=?,Age=?,BirthDate=? WHERE Id=?`, user.Name, user.Age, user.BirthDate, Id)
		if err != nil {
			return
		}
		return

	case http.MethodDelete:
		u := strings.Split(r.URL.Path, "/user/")
		Id, err := strconv.Atoi(u[len(u)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, err = db.Query(`DELETE FROM user where Id=?`, Id)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}

	}

}

func main() {

	http.HandleFunc("/users/", ShowUser)
	http.HandleFunc("/user/", ShowUserWithId)
	http.ListenAndServe(":5000", nil)
}
