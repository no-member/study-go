package main

import "fmt"

func main() {
	var num1_02 int = 3
	var num2_02 float32 = 2.2

	//fmt.Print(num1_02 + num2_02)

	fmt.Println(float32(num1_02) + num2_02)
	fmt.Print(num1_02 + int(num2_02))
}
