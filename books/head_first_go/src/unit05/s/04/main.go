package main

import "fmt"

func main() {
	notes := [3]string{"do", "re", "mi"}
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(notes)
	fmt.Println(primes)
	fmt.Println()

	// %#v의 형식을 지정하면 전달된 배열은 Go 배열 리터럴의 형태로 출력됩니다.
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)
}
