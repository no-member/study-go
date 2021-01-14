package main

import "fmt"

// 포인터 사용법
func main() {
	var value int = 2
	var pointer *int = &value
	fmt.Println(pointer)  // 포인터 주소가 출력된다.
	fmt.Println(*pointer) // 포인터 주소에 있는 값이 출력된다.
}
