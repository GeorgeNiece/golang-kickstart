package main

import "fmt"

func main() {

	funcA()
	fmt.Println("finishing...")

}

func funcC() {
	fmt.Println("funcC")
	panic("Error!")
}

func funcB() {
	fmt.Println("funcB")
	defer defer1()
	funcC()
}

func funcA() {
	fmt.Println("funcA")
	defer defer1()
	defer defer2()
	funcB()
}

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
	/*s:=recover()
	fmt.Println("recovered from ", s)
	*/
}
