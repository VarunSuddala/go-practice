package main

import "fmt"

func urlShortener() (func(string) string, func(string) string) {
	count := 0
	db := make(map[string]string)

	shortener := func(orginal string) string {
		count++

		short := fmt.Sprintf("url%d", count)
		db[short] = orginal

		return short
	}
	resolve := func(url string) string {
		return db[url]
	}
	return shortener, resolve
}
