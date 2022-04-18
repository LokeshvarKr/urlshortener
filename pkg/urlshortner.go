package pkg

import "strings"

//baseURL is constant to add prefix in short URL
const baseURL = "http://infra.cd/"

// ProcessRequestURLAndGetResponseURL takes RequestURLShort data and process and finally return ResponseURLShort
func ProcessRequestURLAndGetResponseURL(requesltURLShort *RequestURL) *ResponseURL {
	url := requesltURLShort.URL
	url = strings.Trim(url, " ")
	shortURL := ""
	if actualURLPresentInDatabase(url) {
		shortURL = actualToShort[url]
	} else {
		shortURL = URLShortener(url)
	}

	// store in databases
	actualToShort[url] = shortURL
	shortToActual[shortURL] = url

	// add baseURL prefix
	shortURL = baseURL + shortURL

	responseURL := ResponseURL{
		ActualURL: url,
		ShortURL:  shortURL,
	}
	return &responseURL
}

// URLShortener takes actual long url and it returns short url for the given url
func URLShortener(url string) string {
	shortURL := randomStringAlgorithm()
	for {
		if !shortURLPresentInDatabase(shortURL) {
			break
		}
		shortURL = randomStringAlgorithm()
	}
	return shortURL
}
