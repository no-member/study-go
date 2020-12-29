package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		fmt.Println("Error 1")
		return
	}

	if a == 2 {
		fmt.Println("Error 2")
		return
	}

	if a == 3 {
		// 에러 1 상황으로 중복 코드
		fmt.Println("Error 1")
		return
	}

	fmt.Println(a)
	fmt.Println("Success")

	return
}
