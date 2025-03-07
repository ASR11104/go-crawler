package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func extractLinks(n *html.Node, rawBaseURL string, result *[]string) {
	parsedBaseURL, _ := url.Parse(rawBaseURL)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				parsedURL, err := url.Parse(attr.Val)
				if err == nil {
					// Convert relative URLs to absolute
					absoluteURL := parsedBaseURL.ResolveReference(parsedURL).String()
					*result = append(*result, absoluteURL)
				}
			}
		}
	}
	// Recursively visit child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c, rawBaseURL, result)
	}
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	result := []string{}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("error parsing html: %w", err)
	}
	extractLinks(doc, rawBaseURL, &result)
	return result, nil
}
