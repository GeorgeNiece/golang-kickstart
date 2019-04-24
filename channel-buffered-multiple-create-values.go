package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("With buffers we are now unblocked")
}

func main() {
	fmt.Println("Create a buffered channel that calculates and run twice rapidly")
	//make the channel, channels have to be created before being used
	valueChannel := make(chan int, 2)  //By changing this to a buffered channel, our send operation, c <- value only blocks within our goroutines should the channel be full. 
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)

	time.Sleep(1000 * time.Millisecond)  //sleep to allow our program to show the output of the buffering fmt output
}