package httpasclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResFormat struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var List []ResFormat

func RequestTo() {
	fmt.Println("start")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/5")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
	fmt.Println(resp.Request.UserAgent(), resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"))

	var rf ResFormat = ResFormat{}

	err = json.Unmarshal(body, &rf)

	fmt.Println("end")

	if err != nil {

		log.Fatalln(err)
	}
	fmt.Println(rf)

	RequestListTo()
}

func RequestListTo() {
	fmt.Println("start")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
	fmt.Println(resp.Request.UserAgent(), resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"))

	err = json.Unmarshal(body, &List)

	fmt.Println("end")
	if err != nil {

		log.Fatalln(err)
	}
	// fmt.Println(List)
}
