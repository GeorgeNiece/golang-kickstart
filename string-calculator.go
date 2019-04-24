package main

import (
	"fmt"
	"strconv"
)

func toString(iN float64) string {
	return strconv.FormatFloat(iN, 'f', -1, 64)
}

func toNumber(n string) float64 {
	iN, err := strconv.ParseFloat(n, 64)
	if err != nil {
		fmt.Println("freak out")
	}
	return iN
}
func main() {
	var FnMap = map[string]func(string, string) string{
		"a": func(a, b string) string {
			return toString(toNumber(a) + toNumber(b))
		},
		"s": func(a, b string) string {
			return toString(toNumber(a) - toNumber(b))
		},
		"m": func(a, b string) string {
			return toString(toNumber(a) * toNumber(b))
		},
		"d": func(a, b string) string {
			if b == "0" {
				panic(a + " can not divided by 0.")
			}
			return toString(toNumber(a) / toNumber(b))
		},
	}

	var op, num1, num2 string
	fmt.Scanln(&op, &num1, &num2)
	fn := FnMap[op]
	fmt.Println(fn(num1, num2))
}
