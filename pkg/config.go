package pkg


// RequestURL is schema for request 
type RequestURL struct {
	URL string `json:"url"`
}


// ResponseURL is schema for response
type ResponseURL struct {
	ActualURL string `json:"actualURL"`
	ShortURL  string `json:"shortURL"`
}
