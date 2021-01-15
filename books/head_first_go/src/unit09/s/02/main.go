package main

import "fmt"

type Literts float64
type Gallons float64

// 기본 타입과 동일하게 연산 가능한 사용자 정의 타입
func main() {
	fmt.Println(Literts(1.2) + Literts(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Literts(2.2) / Literts(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Literts(1.2) < Literts(3.4))
	fmt.Println(Literts(1.2) > Literts(3.4))
}
