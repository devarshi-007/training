package handlinghttprequests

import (
	"net/http"
)

func RequestRouteRequest() {
	http.Handle("/About", &ourHandler{Message: "this is about page"})
	http.HandleFunc("/Help", headerTitle(5))
	http.Handle("/", &ourHandler{Message: "this is home page"})
}

func StartListening() {
	http.ListenAndServe(":5000", nil)
}

func ManageHandler() {
	RequestRouteRequest()
	StartListening()
}

type ourHandler struct {
	Message string
}

func (f *ourHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func headerTitle(k int) func(w http.ResponseWriter, r *http.Request) {

	msg := "this is something which you can in help page"

	switch k {
	case 1:
		msg = "<h1>" + msg + "</h1>"
	case 2:
		msg = "<h2>" + msg + "</h2>"
	case 3:
		msg = "<h3>" + msg + "</h3>"
	case 4:
		msg = "<h4>" + msg + "</h4>"
	case 5:
		msg = "<h5>" + msg + "</h5>"
	case 6:
		msg = "<h6>" + msg + "</h6>"
	default:
		msg = "<p>" + msg + "</p>"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(msg))
	}

}

// m, _ := time.Parse("2-Jan-2006", "5-10-2023")
// 	println(m)
