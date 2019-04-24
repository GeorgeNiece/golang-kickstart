package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverse_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")

}
func reverse_words_delimiters(s string) string {
	
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	newwords:= strings.FieldsFunc(s, f)
	//fmt.Println(newwords)
	for i, j := 0, len(newwords)-1; i<( len(newwords)) ; i, j = i+1, j-1 {
		fmt.Println(i,j, newwords[i], newwords[j])
		s = strings.Replace(s, newwords[i],newwords[j],1)
		
	}
	return s
}

func main() {
	fmt.Println(reverse("one two three"))
	fmt.Println(reverse_words("one two three"))
	fmt.Println(reverse_words_delimiters("Hello.  This is a good day!") //day. good a is This Hello!
}