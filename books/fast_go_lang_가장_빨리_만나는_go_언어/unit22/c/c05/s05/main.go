package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 부분 슬라이스를 만들면서 용략도 같이 지정하기
	b := a[0:6:8]

	fmt.Println(len(b), cap(b))
	fmt.Println(b)
}
