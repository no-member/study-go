package main

import "fmt"

// len 함수로 배열 길이 구하기
func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(len(primes))
}
