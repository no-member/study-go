package main

import "fmt"

// 배열의 경우의 복사
func main() {
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
