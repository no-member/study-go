package main

import "fmt"

// slice 선언하기
func main() {
	var a []int

	// make를 이용하면 slice의 공간 할당 가능
	// make([]자료형, 길이)
	var b []int = make([]int, 5)
	var c = make([]int, 5)
	d := make([]int, 5)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
