package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//program to start with a number and then guess a single time with random if that is the number
func main() {
	var s string
	var count int = 0
	fmt.Println("\nWhat number should we guess !!")
	fmt.Scanln(&s)
	number, _ := strconv.Atoi(s)
	/*if err == nil {
		fmt.Println("No error in number entry")
		os.Exit(0)
	}
	*/
	for {
		count++
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v is too small.\n", n)
		} else if n > number {
			fmt.Printf("%v is too big.\n", n)
		} else {
			fmt.Printf("We got it, in %v guesses! %v\n", count, n)
			break
		}
	}
}
