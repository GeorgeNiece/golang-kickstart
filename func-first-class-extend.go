package main

import (
  "fmt"
  "time"
)

func helloWorldFunc() {
  fmt.Println("Hello World")
  time.Sleep(1 * time.Second)
}

// abstractionFunction example takes a function 
// as a parameter 
func abstractionFunc(a func()) {
	// Receive function and then immediately call
  a()
}

func main() {
  fmt.Printf("Type: %T\n", helloWorldFunc)
  // here we call our abstractionFunc function
  // passing in our helloWorldFunc
	abstractionFunc(helloWorldFunc)
}