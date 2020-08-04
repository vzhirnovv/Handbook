package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words = map[string]int{}
	fields := strings.Fields(s)
	for word := range fields {
		_, ok := words[fields[word]]
		if ok {
			words[fields[word]]++
		} else {
			words[fields[word]] = 1
		}
	}
	return words
}

func main() {
	wc.Test(WordCount)
}