package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("URL parsing failed with error: %w", err)
	}
	host := parsedURL.Host
	path := parsedURL.Path
	if len(path) > 0 && path[len(path)-1] != '/' {
		path = path + "/"
	}
	if path == "" {
		path = "/"
	}
	url := host + path
	return url, nil
}
