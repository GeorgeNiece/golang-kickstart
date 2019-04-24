package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func fib(n int) *big.Int {
	fn := make(map[int]*big.Int)

	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	startTime := time.Now()
	startFibo,e := strconv.Atoi(os.Args[1])
	if e != nil {
		panic(e)
	}
	endFibo,e := strconv.Atoi(os.Args[2])
	if e != nil {
		panic(e)
	}
	for i :=startFibo; i<= endFibo; i++ {
 		fmt.Printf("Fibonacci for %d is %s\n\n", i, fib(i))
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("time to calculate is %s \n", elapsedTime)

}