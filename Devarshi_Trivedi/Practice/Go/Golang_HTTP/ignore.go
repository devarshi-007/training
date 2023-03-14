package main

// type fooHandler struct {
// 	Message string
// }

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(f.Message))
// }

// func main() {
// 	http.Handle("/foo", &fooHandler{Message: "hello world"})
// }

//or

// func main() {
// 	foo := func(w http.ResponseWriter, _ *http.Response) {
// 		w.Write([]byte("hello world"))
// 	}
// 	http.HandleFunc("/foo", foo)
// }
