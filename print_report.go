package main

import "sort"

func sortPages(pages map[string]int) []string {
	keys := []string{}
	for key := range pages {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return pages[keys[i]] > pages[keys[j]]
	})

	return keys
}
