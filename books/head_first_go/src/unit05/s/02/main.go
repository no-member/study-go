package main

import "fmt"

// 배열 리터럴 사용하기
func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	fmt.Println()

	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2], primes[4])
}
