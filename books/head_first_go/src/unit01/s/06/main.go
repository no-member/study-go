package main

import "fmt"

// 제로 값
func main() {
	var myInt int
	var myFloat float64

	fmt.Println("myInt: ", myInt)     // 0
	fmt.Println("myFloat: ", myFloat) // 0

	var myString string
	var myBool bool

	fmt.Println("myString: ", myString) // ""
	fmt.Println("myBool: ", myBool)     // false
}
