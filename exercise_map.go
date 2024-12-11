package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	strArr := strings.Fields(s)
	for i := range strArr {
		words[strArr[i]] += 1
	}
	
	return words
}

func main() {
	wc.Test(WordCount)
}

