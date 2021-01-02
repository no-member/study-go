package main

import "fmt"

// slice를 생성하면서 값을 추기화 하기
func main() {
	// 슬라이스를 생성하면서 값을 초기화
	a := []int{32, 29, 78, 16, 81}

	// 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	b := []int{
		32,
		29,
		78,
		16,
		81,
	}

	fmt.Println(a)
	fmt.Println(b)
}
