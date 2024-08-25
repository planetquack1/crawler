package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("too few arguments provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(1)
		fmt.Println(fmt.Errorf("cannot convert max concurrency to an int"))
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		os.Exit(1)
		fmt.Println(fmt.Errorf("cannot convert max pages to an int"))
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)

	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }
}
