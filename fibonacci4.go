package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

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
	for _, s := range os.Args[1:] {
		n, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		fmt.Printf("%s\n", fib(n))
	}

}