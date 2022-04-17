package pkg

type RequestURL struct {
	URL string `json:"url"`
}

type ResponseURL struct {
	ActualURL string `json:"actualURL"`
	ShortURL  string `json:"shortURL"`
}
