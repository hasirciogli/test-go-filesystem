package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	server
}
func ProductsHandler(writer http.ResponseWriter, request *http.Request) {

}
func ArticlesHandler(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
