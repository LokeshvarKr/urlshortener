package main

import (
	"net/http"

	"github.com/LokeshvarKr/urlshortener/pkg"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	pkg.RegisterRouters(r)
	http.ListenAndServe(":8081", r)
}
