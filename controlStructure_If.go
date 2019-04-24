package main

import "fmt"

func main() {
	x := basicOne(-21)
	fmt.Println(x)

	// Type assertion inside if
	var val interface{}
	val = "foo"
	if str, ok := val.(string); ok {
		fmt.Println(str)
	}
}
func basicOne(x int) int {
	//Basic one
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func basicTwo(b, c int) int {
	// You can put one statement before the condition
	if a := b + c; a < 42 {
		return a
	} else {
		return a - 42
	}
}
