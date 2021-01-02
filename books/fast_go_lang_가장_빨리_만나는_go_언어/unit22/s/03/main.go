package main

import "fmt"

func main() {
	// make([]타입, 길이, 용량)
	a := make([]int, 5, 10)

	// slice의 길이를 구하는 len
	fmt.Println(len(a))

	// slice의 용량을 구하는 cap
	fmt.Println(cap(a))
}
