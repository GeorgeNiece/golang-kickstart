package main

import "fmt"
import "math/big"

func main() {
        fmt.Println("Generate a factorial without input")

    //n := big.NewInt(40)
    r := big.NewInt(12)

    fmt.Println(factorial(r))

}

func factorial(n *big.Int) (result *big.Int) {
    //fmt.Println("n = ", n)
    b := big.NewInt(0)
    c := big.NewInt(1)

    if n.Cmp(b) == -1 {
        result = big.NewInt(1)
    }
    if n.Cmp(b) == 0 {
        result = big.NewInt(1)
    } else {
        return n.Mul(n,factorial(n.Sub(n,c)));
        //fmt.Println("n = ", n)
        //result = n.Mul(n, factorial(n.Sub(n, c)))
    }
    return result
}