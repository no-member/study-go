package main

import "fmt"

func main() {
	// ...을 사용하면 초기화 요소의 개수를 생략 가능
	a := [...]int{32, 29, 78, 16, 81}

	b := [...]string{"Maria", "Andrew", "Jhon"}

	fmt.Println(a)
	fmt.Println(b)
}
