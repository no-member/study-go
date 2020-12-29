package main

import "fmt"

func main() {
	i := 3

	switch i {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		// 마지막 case에는 fallthrough를 사용할 수 없음
		fmt.Println("0 이상")
	}
}
