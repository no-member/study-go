package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	// go에서 *pointer.myField라고 표기를 할경우 myField의 포인터를 가져온다고
	// 해석하기 때문에 오류가 발생(myField는 포인터 타입이 아니기 때문에)
	// fmt.Println(*pointer.myField)

	fmt.Println((*pointer).myField)
	// 위의 방식으로 구조체의 필드에 접근 할 수도 있지만 굉장히 번거롭기 때문에
	// 도트 연산자를 이용해서 직접 필드에 접근 가능
	fmt.Println(pointer.myField)

	pointer.myField = 9
	fmt.Println(pointer.myField)
}
