package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("First")
	go func() {
		fmt.Println("???")
	}()
	fmt.Println("Is this 2nd or not")
	if len(os.Args) > 1 {
		fmt.Scanln()
	}
	//fmt.Println(math.Sin(3.1416))
	fmt.Println("done")
}
