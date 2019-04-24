package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	i := 1
	for range time.Tick(time.Second) {
		// do it once a sec
			if i > 10 {
				os.Exit(0)
			}
			fmt.Println(time.Now())
			i++
	
	}
}