package main

import "fmt"

var meterPerLiter float64 // 패키지 수준에서 변수를 선언한 경우 어디에서든 접근 가능

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / meterPerLiter
}

func main() {
	meterPerLiter = 10.0
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
}
