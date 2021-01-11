package main

import "fmt"

// 제로 값
func main() {
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[8], boolSlice[0])
}
