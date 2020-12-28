package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println()

	// for을 while과 같은 형태로 작성하기
	i := 0
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()
	fmt.Println(i)
}
