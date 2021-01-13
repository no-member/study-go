package main

import "fmt"

// 변수에 구조체 선언하기
func main() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)
}
