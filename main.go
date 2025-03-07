package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	url := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("maxConcurrency should be a string")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("maxPages should be a string")
		os.Exit(1)
	}
	cnf := configure(url, maxConcurrency, maxPages)
	cnf.wg.Add(1)
	go cnf.crawlPage(url)
	cnf.wg.Wait()
	cnf.printReport()
	os.Exit(0)
}
