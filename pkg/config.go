package pkg

type RequestURLShort struct {
	ActualURL string `json:"actualURL"`
}

type ResponseURLShort struct {
	ActualURL string `json:"actualURL"`
	ShortURL  string `json:"shortURL"`
}
