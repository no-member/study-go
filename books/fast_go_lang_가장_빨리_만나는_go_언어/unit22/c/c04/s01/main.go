package main

import "fmt"

// slice 용량 늘리기
func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(len(a), cap(a))
	fmt.Println()

	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))
}
