package main

import (
	"fmt"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func configure(url string, concurrency int, maxPages int) *config {
	return &config{
		pages:              map[string]int{},
		baseURL:            url,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, concurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) maxPageExceeded() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	visitedPages := len(cfg.pages)
	return visitedPages >= cfg.maxPages
}

func (cfg *config) printReport() {
	sortedPages := sortPages(cfg.pages)
	fmt.Println("=============================")
	fmt.Println("REPORT for", cfg.baseURL)
	fmt.Println("=============================")
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", cfg.pages[page], page)
	}

}
