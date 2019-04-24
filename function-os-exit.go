package main

import "fmt"
import "os"

func main() {
	
	defer fmt.Println("deferred to end of run")
	fmt.Println("running test on os.Exit")
	os.Exit(1)

}

