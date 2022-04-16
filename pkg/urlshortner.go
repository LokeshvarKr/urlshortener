package pkg

//baseURL is constant to add prefix in short URL
const baseURL = "http://infra.cd/"

// ProcessRequestURLShortData takes RequestURLShort data and process and finally return ResponseURLShort
func ProcessRequestURLShortData(requesltURLShort *RequestURLShort) *ResponseURLShort {
	url := requesltURLShort.ActualURL
	shortURL := URLShortener(url)
	responseURLShort := ResponseURLShort{
		ActualURL: url,
		ShortURL:  shortURL,
	}
	return &responseURLShort
}

// URLShortener takes actual long url and it returns short url for the given url
func URLShortener(url string) string {
	shortURL := randomStringAlgorithm()
	for {
		if _, ok := database[shortURL]; !ok {
			break
		}
		shortURL = randomStringAlgorithm()
	}

	// store url corresponding to shortURL
	database[shortURL] = url

	return baseURL + shortURL
}
