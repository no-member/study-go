package main

import "fmt"

// 슬라이스 리터럴
func main() {

	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	fmt.Println()

	primes := []int{
		2,
		3,
		5,
	}
	fmt.Println(primes[0], primes[1], primes[2])
}
