// go run short-circuit.go
package main
import "fmt"

func AFunc() bool {
	return true
}

func BFunc() bool {
	return false
}

func val1() bool{
    fmt.Println("val1 gets called")
    return true
}

func val2() bool{
    fmt.Println("val2 gets called")
    return false
}

func main(){
	fmt.Println("Condition 1: val1() && val2()")
    if val1() && val2() {
        fmt.Println("Shouldn't print")
    }
	fmt.Println("Condition 2: val2() && val1()")
    if val2() && val1() {
        fmt.Println("Shouldn't print")
    }
	fmt.Println("Condition 3: val1() || val2()")
    if val1() || val2(){
        fmt.Println("The boolean expression is true")
    }
	fmt.Println("Condition 4: val2() || val1()")
    if val2() || val1(){
        fmt.Println("The boolean expression is true")
    }
	var bool1 = AFunc() || BFunc()
	var bool2 = AFunc() || true
	var bool3 = false && BFunc()
 	var bool4 = AFunc()
	var bool5 = BFunc() || AFunc() || false
	fmt.Println(bool1, bool2, bool3, bool4, bool5)
}