package main

import "fmt"
import "sort"

func main() {

	myslice := []int{11,2,213,23,4,66}
	sort.Ints(myslice)
	fmt.Println(myslice)
	fmt.Println(sort.IntsAreSorted(myslice))


}