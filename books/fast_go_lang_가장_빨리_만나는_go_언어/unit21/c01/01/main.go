package main

import "fmt"

// len을 이용한 배열 순회
func main() {
	a := [5]int{32, 29, 78, 16, 81}

	// 배열의 길이를 구하는 len 함수
	fmt.Println(len(a))
	fmt.Println()

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
