package main

import "fmt"

// 부분 슬라이스 사용하기
func main() {
	a := []int{1, 2, 3, 4, 5}

	b := a[0:5]

	fmt.Println(a)
	fmt.Println(b)
}
