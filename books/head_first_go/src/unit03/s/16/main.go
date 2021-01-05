package main

import "fmt"

func main() {
	amount := 6
	double(amount)      // 12 <= 함수에 인자를 전달
	fmt.Println(amount) // 6 <= 원래 값을 출력
}

func double(number int) {
	number *= 2
	fmt.Println(number) // <= 원래의 값이 아닌 복사된 값을 변경
}
