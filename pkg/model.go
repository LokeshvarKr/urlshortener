package pkg

import (
	"fmt"
	"strings"
)

// here maps are in-memory database
var shortToActual = make(map[string]string)
var actualToShort = make(map[string]string)

func RetriveActualURLAndGetResponseURL(requestURL *RequestURL) (*ResponseURL, error) {
	url := requestURL.URL
	modifiedURL := strings.Trim(url, " ")
	modifiedURL = strings.Replace(modifiedURL, baseURL, "", 1)

	if shortURLPresentInDatabase(modifiedURL) {
		actualURL := shortToActual[modifiedURL]
		responseURL := ResponseURL{
			ActualURL: actualURL,
			ShortURL:  requestURL.URL,
		}
		return &responseURL, nil
	}
	return nil, fmt.Errorf("no such url found in db")
}



func actualURLPresentInDatabase(actual string) bool {
	if _, ok := actualToShort[actual]; ok {
		return true
	}
	return false
}

func shortURLPresentInDatabase(shortURL string) bool {
	if _, ok := shortToActual[shortURL]; ok {
		return true
	}
	return false
}
