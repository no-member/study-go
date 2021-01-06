package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i, v := range notes {
		fmt.Println(i, v)
	}
	fmt.Println()

	// value 값을 사용 할 필요가 없다.
	for i := range notes {
		fmt.Println(i)
	}
	fmt.Println()

	// index를 사용하고 싶지 않을 경우에는 빈 식별자를 사용하자
	for _, v := range notes {
		fmt.Println(v)
	}
	fmt.Println()
}
