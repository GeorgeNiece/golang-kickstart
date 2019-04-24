package main

import (
	"fmt"
)

var a = []string{"Hello", "Hola", "Bon Jour", "Ciao"}
func main() {
	fmt.Println(a)
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	fmt.Println(a)
}
