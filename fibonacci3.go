package main

import (
	"fmt"
	"strconv"
)

func fibonacci() func(int) int {
	arr := []int{0, 1}
	return func(num int) int {
		arr = append(arr, arr[num+1]+arr[num])
		return arr[num]
	}
}

func main() {
	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn(i))
	}
}