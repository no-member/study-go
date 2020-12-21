package main

import "fmt"

func main() {
	var num1 uint16 = 10
	var num2 uint32 = 80000

	fmt.Println(num1 + uint16(num2))
}
