package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	r.HandleFunc("/home", HomePage).Methods("GET")
	r.HandleFunc("/shorturl", ShortURL).Methods("POST")
	r.HandleFunc("/actualurl", ActualURL).Methods("POST")
}

// ShortURL Short given URL (POST /shorturl)
func ShortURL(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	// read Body
	requestURL := &RequestURL{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(requestURL)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !checkValidRequestURL(requestURL.URL) {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	// process RequestURLShort data and get ResponseURLShort
	responseURL := ProcessRequestURLAndGetResponseURL(requestURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// encode request into json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// ActualURL returns atual url to given short URL (POST /actualurl)
func ActualURL(w http.ResponseWriter, r *http.Request) {

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	// read Body
	requestURL := &RequestURL{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(requestURL)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrive actual url from db and get ResponseURLShort
	responseURL, err := RetriveActualURLAndGetResponseURL(requestURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// encode request into json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("URL shortner home page"))
}
