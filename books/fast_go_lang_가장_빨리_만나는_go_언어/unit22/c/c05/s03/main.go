package main

import "fmt"

// 부분 슬라이스 사용하기
func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[:])
	fmt.Println(a[0:])
	fmt.Println(a[0:5])
	fmt.Println(a[0:len(a)])
	fmt.Println()

	fmt.Println(a[3:])
	fmt.Println(a[:3])
	fmt.Println(a[1:3])
}
