package main

import "fmt"

// 배열에서 부분 슬라이스 사용하기
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:2]

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println()
	b[0] = 9

	// 부분 슬라이스를 한 경우 참조이기 때문에 요소를 변화시키면
	// 	배열의 요소도 변화
	fmt.Println("a", a)
	fmt.Println("b", b)
}
