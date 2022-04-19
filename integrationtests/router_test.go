package pkg

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	// requestURL := pkg.RequestURL{URL: "http://asjfdlsa.com/sdlfjsldfjsd/sdlkjfa"}
	req, err := http.NewRequest("GET", "http://localhost:8081/home", nil)

	assert.Nil(t, err)
	assert.NotNil(t, req)

	c := http.Client{}

	resp, err := c.Do(req)
	assert.NotNil(t, resp)
}

func TestShortURL(t *testing.T) {
	// requestURL := pkg.RequestURL{URL: "http://asjfdlsa.com/sdlfjsldfjsd/sdlkjfa"}
	req, err := http.NewRequest("POST", "http://localhost:8081/shorturl", nil)

	assert.Nil(t, err)
	assert.NotNil(t, req)

	c := http.Client{}

	resp, err := c.Do(req)
	assert.NotNil(t, resp)
}

func TestActualURL(t *testing.T) {
	// requestURL := pkg.RequestURL{URL: "http://asjfdlsa.com/sdlfjsldfjsd/sdlkjfa"}
	req, err := http.NewRequest("POST", "http://localhost:8081/actualurl", nil)

	assert.Nil(t, err)
	assert.NotNil(t, req)

	c := http.Client{}

	resp, err := c.Do(req)
	assert.NotNil(t, resp)
}
