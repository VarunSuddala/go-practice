package main

import (
	"math/rand"
	"time"
)

func urlShortener() (func(string) string, func(string) string) {
	rand.Seed(time.Now().UnixNano())

	shortToOriginal := make(map[string]string)
	originalToShort := make(map[string]string)

	shortener := func(original string) string {
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		if short, ok := originalToShort[original]; ok {
			return short
		}
		for {
			code := make([]byte, 6)
			for i := range code {
				code[i] = chars[rand.Intn(len(chars))]
			}
			short := "url" + string(code)

			if _, exists := shortToOriginal[short]; !exists {
				originalToShort[original] = short
				shortToOriginal[short] = original
				return short
			}
		}
	}
	resolve := func(url string) string {
		return shortToOriginal[url]
	}
	return shortener, resolve
}
