package main

import (
	"fmt"
	"reflect"
)

func main() {
	funcA(5,"cool")
	funcB(5, true, true, false)
}

func funcA(a int, b string){
	fmt.Println(a,b)
}

func funcB(a int, b ...bool) {
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}

