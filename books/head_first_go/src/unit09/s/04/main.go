package main

import "fmt"

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

type MyType string

// 메서드 사용하기
func main() {
	value := MyType("a MyType value")
	value.sayHi()
}
