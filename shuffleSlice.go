package main

import (
	"fmt"
	"math/rand"
)

var a = []string{"Hello", "Hola", "Bon Jour", "Ciao"}
func main() {
	fmt.Println(a)
	for i := len(a) - 1; i > 0; i-- {
    		j := rand.Intn(i + 1)
    		a[i], a[j] = a[j], a[i]
	}	
	fmt.Println(a)
}
