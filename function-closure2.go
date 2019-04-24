package main

import (
	"fmt"
	
)

	func aFunc() func() {
		a:= 5
		return func() {
			a++
			fmt.Println(a)
		}
	}


func main() {
	f:=aFunc()
	f()
}

