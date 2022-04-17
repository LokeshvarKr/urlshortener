package pkg

import (
	"math/rand"
	"net/url"
	"time"
)

// min and max are constants to generate random integer in range(min,max)
const (
	min int = 1000000000
	max int = 9999999999
)

//  a-zA-Z1-9 ----> total 62 char we can consider but
// 0(zero) and capital O looks smiliar to human
// 1(one) and small l lokks smiliar to human
// So not considering 4 characters  0(zero), capital O, 1(one), small l  ----> 0O1l
// only 58 charaters will be used in short generated url
var lookUp []rune = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ23456789")

// use LookUp[] to convert random integer into base58 format
func base58(n int) string {
	output := ""
	for n > 0 {
		x := n % 58
		n = n / 58
		output += string(lookUp[x])
	}

	// reverse string
	temp := []rune(output)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[j], temp[i] = temp[i], temp[j]
	}
	output = string(temp)
	return output
}

func randomStringAlgorithm() string {
	// in first call only, seed with crruent unix time to generate random unmber
	t := time.Now().UnixNano()
	rand.Seed(t)

	randInt := rand.Intn(max-min+1) + min
	output := base58(randInt)
	return output

}

//CheckValidRequestURL validates if given url domain name is our domain name
// if given url domain name is our domain then return false (means says invalid url)
func CheckValidRequestURL(actualURL string) bool {
	u, err := url.ParseRequestURI(actualURL)
	if err != nil || u.Host == "" || u.Host == baseURL {
		return false
	}
	return true
}
