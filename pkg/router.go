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
	requestURLShort := &RequestURLShort{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(requestURLShort)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// process RequestURLShort data and get ResponseURLShort
	responseURLShort := ProcessRequestURLShortData(requestURLShort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// encode request into json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseURLShort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// ShortURL Short given URL (POST /actualurl)
func ActualURL(w http.ResponseWriter, r *http.Request) {

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("URL shortner home page"))
}
