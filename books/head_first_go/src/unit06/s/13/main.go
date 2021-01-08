package main

import "fmt"

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

// 가변 인자
func main() {
	severalInts(1)
	severalInts(1, 2, 3)
	severalInts() // 인자가 없는 경우 빈 슬라이스를 인자로 받는다.
}
