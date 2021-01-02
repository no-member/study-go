package main

import "fmt"

// append를 이용해 슬라이스에 값 추가
func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println()

	a = append(a, 4, 5, 6)

	fmt.Println(a)
}
