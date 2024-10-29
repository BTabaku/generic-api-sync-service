// internal/api/fetch.go
package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

// FetchData retrieves data from the specified URL
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}
	return data, nil
}
