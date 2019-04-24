package main

import (
	"fmt"
	
)

func main() {
	a:= 5
	b:=10
	x:= func() int {
		return a*b
	}

	fmt.Println(x())
	fmt.Println(func() int { 
		return 42
	}())
}

