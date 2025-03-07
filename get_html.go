package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	fmt.Println("Getting Html")
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error while getting the page: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "", fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}
	if err != nil {
		return "", fmt.Errorf("error while reading the page: %w", err)
	}
	return string(body), nil
}
