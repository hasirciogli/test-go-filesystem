package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "video/mp4")
	writer.Write([]byte("here is a string dude"))
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

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
