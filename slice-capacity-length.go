package main

import "fmt"

func main() {
	//var2 := make(map[string]int, 9)
	//fmt.Println(cap(var2))

	var3 := make([]int, 5, 10)
	fmt.Println(cap(var3))

	s := make([]int, 3, 10)
    	fmt.Println("capacity:", cap(s))
	fmt.Println("length:", len(s))
	fmt.Println(s)

	t := make([]int, 3)
    	fmt.Println("capacity:", cap(t))
	fmt.Println("length:", len(t))
	fmt.Println(s)
}
