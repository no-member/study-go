package main

import "fmt"

func facItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(facItr2(5))
}
