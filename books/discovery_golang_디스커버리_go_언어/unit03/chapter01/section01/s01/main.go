package main

import "fmt"

func main() {
	for i, r := range "가나다" {
		// i에는 바이트 위치
		// r에는 rune이 들어감
		fmt.Println(i, r)
	}
	fmt.Println()

	fmt.Println(len("가나다"))
}
