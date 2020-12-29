package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)
	fmt.Println(a == 9.0) // => false
	fmt.Println()

	const epsilon = 1e-14 // go 언어의 입실론
	fmt.Println(math.Abs(a-9.0) <= epsilon)
}
