package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func countOneToZero() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

// 두개 이상의 반환 값 중에서 사용하고 싶지 않는 값을 생략 가능
func main() {
	_, diff2 := SumAndDiff(10, 2)
	fmt.Println(diff2)

	a, _, c, _, e := countOneToZero()
	fmt.Println(a, c, e)
}
