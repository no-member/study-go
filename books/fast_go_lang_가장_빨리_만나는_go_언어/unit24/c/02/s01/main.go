package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

// 반환 값을 두개 이상 가질 수 있는 Go
func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)
}
