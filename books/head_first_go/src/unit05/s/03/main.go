package main

import "fmt"

// 배열 리터럴과 단축 변수 선언
func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println(notes[0])
	fmt.Println(primes[0])

	// 줄 바꿈 문자가 오는 경우에는 쉼표를 붙여 줘야 합니다.
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}

	fmt.Println(text)
}
