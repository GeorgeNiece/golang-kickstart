package main

import (
  "fmt"
  "time"
)

func helloWorldFunc() {
  fmt.Println("Hello First Class Function World")
  time.Sleep(1 * time.Second)
}

func main() {
  fmt.Printf("Type: %T\n", helloWorldFunc)
}