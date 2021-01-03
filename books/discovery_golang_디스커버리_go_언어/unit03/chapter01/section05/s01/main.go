package main

import "fmt"

func main() {
	s := "abc"
	ps := &s
	// 간단한 문자열 이어 붙일 때 사용할 수 있는 +=
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	fmt.Println()

	n := 123
	// 문자열이 아닌 다른 데이터도 이어 붙일 수 있는 Sprint
	nStr := fmt.Sprint(n, 456)
	fmt.Println(nStr)
}
