package main

import "fmt"

func main() {

	var m4=map[string] int {
		"a":1,
		"b":2
	}
	m4["c"] = 3
	fmt.Println(m4)
}