package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (ret map[string]int) {
	ret = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count := ret[word]
		ret[word] = count + 1
	}
	return
}

func main() {
	wc.Test(WordCount)
}
