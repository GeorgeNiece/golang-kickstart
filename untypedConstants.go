package main

import "fmt"

func main() {
	//const a int = 5
	const a = 5
	const b = 5
	var c float32 = a
	var d float32 = b
	fmt.Printf("%f\n%f\n", c, d)
	fmt.Printf("%d\n%d\n", c, d)
	fmt.Printf("%e\n%e\n", c, d)
	fmt.Printf("%d\n%d\n", a, b)

}