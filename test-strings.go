package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse_the_words("How are you"))
}

func reverse_the_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
