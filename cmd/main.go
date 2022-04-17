package main

import (
	"net/http"

	"github.com/LokeshvarKr/urlshortener/pkg"
	"github.com/gorilla/mux"
)

func main() {

	// create a router from gorilla/mux 
	r := mux.NewRouter()

	// register all the enpoints 
	pkg.RegisterRouters(r)

	// start listing and serving http request 
	http.ListenAndServe(":8081", r)
}
