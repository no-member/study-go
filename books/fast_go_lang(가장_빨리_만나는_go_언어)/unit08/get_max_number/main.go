package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxInt16
	var num3 uint32 = math.MaxInt32
	var num4 uint64 = math.MaxInt64

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
