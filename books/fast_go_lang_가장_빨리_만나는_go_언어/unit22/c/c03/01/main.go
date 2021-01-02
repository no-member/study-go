package main

import "fmt"

// slice의 깊은 복사 copy
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3)

	copy(b, a)
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
