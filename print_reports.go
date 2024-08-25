package main

import (
	"fmt"
	"sort"
)

type kv struct {
	Key   string
	Value int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Value, page.Key)
	}
}

func sortPages(pages map[string]int) []kv {

	// Create a slice to hold the keys of the map
	var sortedPages []kv

	// Fill the slice with the key-value pairs from the map
	for k, v := range pages {
		sortedPages = append(sortedPages, kv{k, v})
	}

	// Sort the slice by the map's values in descending order
	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Value > sortedPages[j].Value
	})

	return sortedPages
}
