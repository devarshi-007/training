package main

import (
	"log"
	"main/user"
	"net/http"
)

//var userList []user.User

// func userDD() {
// 	filename := "data/userD.json"
// 	_, err := os.Stat(filename)
// 	if err != nil {
// 		return
// 	}
// 	var userTempList []user.User
// 	data, err := ioutil.ReadFile(filename)
// 	fmt.Println(data)
// 	if err != nil {
// 		fmt.Println("error 1")
// 		return
// 	}
// 	err = json.Unmarshal([]byte(data), &userTempList)
// 	if err != nil {
// 		fmt.Println("error 2", err)
// 		return
// 	}
// 	data, err = json.Marshal(userTempList)
// 	if err != nil {
// 		fmt.Println("error 3")
// 		return
// 	}
// 	fmt.Println(data)
// }

const basePAth = "/api"

func main() {
	// userDD()
	user.SetUpRoutes(basePAth)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
