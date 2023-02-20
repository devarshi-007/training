package main

// import (
// 	"net/http"
// )

// type Message struct {
// 	MessageName string
// 	MessageId   int
// }

// func (m *Message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(m.MessageName))
// 	// w.Write([]byte(string(m.MessageId)))
// }

// func main() {
// 	http.Handle("/mes", &Message{"Hi", 0})
// 	http.ListenAndServe(":4444", nil)
// }
