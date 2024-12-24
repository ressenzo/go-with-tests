package concurrency

import "net/http"

// This is the func to be passed and check website,
// but it is not used
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	return response.StatusCode == http.StatusOK
}
