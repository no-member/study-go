package main

import "fmt"

// len 함수를 이용해 슬라이스 길이를 구하기
func main() {
	notes := make([]string, 7)
	primes := make([]int, 5)
	fmt.Println(len(notes))
	fmt.Println(len(primes))
}
