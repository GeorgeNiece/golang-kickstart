package main 

import (
	"fmt" 
	"unsafe"
)

func main () {

var i int = 1
var ui uint = 10
var f32 float32 = 36.9
fmt.Printf("Size of i is: %d\n", unsafe.Sizeof(i))
fmt.Printf("Size of i is: %d\n", unsafe.Sizeof(ui))
fmt.Printf("Size of i is: %d\n", unsafe.Sizeof(f32))
}