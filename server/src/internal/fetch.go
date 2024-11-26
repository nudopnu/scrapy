package internal

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
)

func Fetch(url string, headers map[string]string) ([]byte, error) {
	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	for header := range headers {
		req.Header.Set(header, headers[header])
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the server sent a gzip-encoded response
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		// Create a new gzip reader if the response is compressed
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
	default:
		// Use the response body directly if it's not compressed
		reader = resp.Body
	}

	// Read and print the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}
