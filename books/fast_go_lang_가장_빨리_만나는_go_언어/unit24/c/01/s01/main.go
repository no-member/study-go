package main

import "fmt"

// 리턴문에 변수명 지정하기
func sum(a int, b int) (r int) {
	return a + b
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
