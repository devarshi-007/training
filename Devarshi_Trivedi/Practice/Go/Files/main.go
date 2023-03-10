package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

const maxUploadSize = 3 * 1024 * 1024

func uploadFile(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ui.html")
	t.Execute(w, nil)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
		return
	}
	file, fHeader, err := r.FormFile("uploadFile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(err)
		w.Write(data)
		return
	}
	defer file.Close()
	// Get and print out file size
	fileSize := fHeader.Size
	fmt.Printf("File size (bytes): %v\n", fileSize)
	// validate file size
	if fileSize > maxUploadSize {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(fileBytes)

}

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":5000", nil)
}
