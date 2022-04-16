package pkg

import "net/http"

func ShortURL(w http.ResponseWriter, r *http.Request) {

}

func ActualURL(w http.ResponseWriter, r *http.Request) {

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("URL shortner home page"))
}
