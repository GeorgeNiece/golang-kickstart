package main

import (
	"fmt"
)

var a = []string{"Hello", "Hola", "Bon Jour", "Ciao"}
func main() {
	fmt.Println(a)
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)
}
