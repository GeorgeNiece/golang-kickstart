//Program to find Factorial of number
package main

import "fmt"

/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.
var i int = 1
var n int

/*     function declaration        */
func factorial(n int) uint64 {
	if n < 0 {
		// cannot use negative numbers in factorial
		panic("Factorial of negative number doesn't exist.")
	} else if n > 50 {
		// not allowing factorial values greater than 50 due to overflow
		panic("Factorial of number greater than 50 not supported")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}

func main() {
	defer func() {
		coverage := recover()
		if coverage != nil {
			fmt.Println("Error encountered and handled:", coverage)
		}
	}()
	fmt.Print("Enter a positive integer between 0 - 50 : ")
	fmt.Scan(&n)
	fmt.Print("Factorial is: ", factorial(n))
}
