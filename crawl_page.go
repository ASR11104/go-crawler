package main

import (
	"fmt"
	"net/url"
)

func (c *config) crawlPage(rawCurrentURL string) error {
	c.concurrencyControl <- struct{}{}
	defer func() {
		c.wg.Done()
		<-c.concurrencyControl
	}()
	if c.maxPageExceeded() {
		return nil
	}
	rawBaseURL := c.baseURL
	fmt.Printf("Crawling %s ...\n", rawCurrentURL)
	baseUrlParse, err1 := url.Parse(rawBaseURL)
	currentUrlParse, err2 := url.Parse(rawCurrentURL)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("error parsing url")
	}
	baseUrlDomain := baseUrlParse.Host
	currentUrlDomain := currentUrlParse.Host
	if baseUrlDomain != currentUrlDomain {
		fmt.Println("Domains not matching")
		return nil
	}

	normalizedCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting normalized url: %v", err)
		return fmt.Errorf("error getting normalized url: %w", err)
	}

	// Lock the pages map during update to avoid deadlock
	if !c.addPageVisit(normalizedCurrentUrl) {
		return nil
	}

	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html body: %v", err)
		return fmt.Errorf("error getting html body: %w", err)
	}
	urls, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting urls from html body: %v", err)
		return fmt.Errorf("error getting urls from html body: %w", err)
	}
	for _, url := range urls {
		c.wg.Add(1)
		go c.crawlPage(url)
	}
	return nil
}
