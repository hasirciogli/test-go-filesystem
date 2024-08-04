package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hasirciogli/test-go-filesystem/server"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(server.CreateServer())
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
