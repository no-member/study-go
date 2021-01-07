package main

import "fmt"

// 가변 인자 사용하기
func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func main() {
	r := sum(1, 2, 3, 4, 5)
	fmt.Println(r)
}
