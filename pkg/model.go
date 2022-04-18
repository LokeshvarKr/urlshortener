package pkg

import (
	"fmt"
	"strings"
)

// Later either sql or no-sql database will be used.

// for now using two maps to store data

// shortToActual stores short URL as key and actual URL as value.
var shortToActual = make(map[string]string)

//actualToShort stores actual URL as key and Short URL as value
var actualToShort = make(map[string]string)


// RetriveActualURLAndGetResponseURL returns ResponseURL if actual URL corresponding to short URL is present in database 
// otherwise it return err 
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

// actualURLPresentInDatabase to check if actual URL is present in database
func actualURLPresentInDatabase(actual string) bool {
	if _, ok := actualToShort[actual]; ok {
		return true
	}
	return false
}

// shortURLPresentInDatabase to check if short URL is present in database
func shortURLPresentInDatabase(shortURL string) bool {
	if _, ok := shortToActual[shortURL]; ok {
		return true
	}
	return false
}
