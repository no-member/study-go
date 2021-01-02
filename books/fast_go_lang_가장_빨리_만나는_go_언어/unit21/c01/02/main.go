package main

import "fmt"

// range를 이용한 배열 순회
func main() {
	a := [5]int{32, 29, 78, 16, 81}

	// index 값과 value값을 출력
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println()

	// index 값만 출력
	for i := range a {
		fmt.Println(i)
	}
	fmt.Println()

	// value 값만 출력
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println()
}
