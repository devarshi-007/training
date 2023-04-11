package database

import "net/http"

// models
type Product struct {
	Id   int
	Name string `json:"product_name"`
}

// serverHttp for ~/About and ~/
func (f *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Name))
}
