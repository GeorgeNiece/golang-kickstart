package main

import (
	"fmt"
	"math"
	"math/big"
)

if true {
	fmt.Println("this works")
}
func main () {
	fmt.Println(mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64)))
}
func mul(x, y *big.Int) *big.Int {

}

func sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
