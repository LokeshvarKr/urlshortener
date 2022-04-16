package pkg

import (
	"github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	r.HandleFunc("/home", HomePage).Methods("GET")
	r.HandleFunc("/shorturl", ShortURL).Methods("POST")
	r.HandleFunc("/actualurl", ActualURL).Methods("POST")
}
