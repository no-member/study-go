package main

import "fmt"

// 배열 복사하기
func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a // 배열을 대입하면 배열 전체가 복사

	fmt.Println("a", a)
	fmt.Println("b", b)
}
