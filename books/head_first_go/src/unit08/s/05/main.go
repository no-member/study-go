package main

import (
	"fmt"
)

type part struct {
	description string
	count       int
}

// 사용자 정의 타입을 함수에서 사용하는 방법
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count", p.count)
}

// 함수를 이용해서 구조체의 값을 반환하는 방법
func minimuOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)
	fmt.Println()

	p := minimuOrder("Hex bolts")
	fmt.Println(p.description, p.count)
}
