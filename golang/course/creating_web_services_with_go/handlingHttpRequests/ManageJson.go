package handlinghttprequests

import (
	"encoding/json"
	"fmt"
	"log"
)

type foo struct {
	Message string `json:"my_message"`
	Name    string `json:"my_name"`
	Age     int    `json:"age"`
}

func genJson() {
	strct := foo{Message: "this is my message", Name: "_name_", Age: 15}
	data, _ := json.Marshal(&strct)
	fmt.Println(string(data))

	err := json.Unmarshal(data, &strct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strct)
}
