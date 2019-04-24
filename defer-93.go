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
	defer defer3()
	funcC()
}

func funcA() {
	defer defer1()
	defer defer2()
	funcB()
}

func defer1() {
	fmt.Println("defer1")
}

func defer3() {
	fmt.Println("defer3")
}

func defer2() {
	fmt.Println("defer2")
	s := recover()
	fmt.Println("recovered from ", s)
}
