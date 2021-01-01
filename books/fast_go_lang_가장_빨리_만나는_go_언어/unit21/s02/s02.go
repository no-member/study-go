package main

import "fmt"

func main() {
	// 배열을 선언하는 여러가지 방법

	// 형식과 초기화를 동시에
	var a [5]int = [5]int{32, 29, 78, 16, 81}

	// 초기화를 할 때 자료형과 길이를 생략
	var b = [5]int{32, 29, 78, 16, 81}

	// 배열을 선언할 때 var 키워드, 자료형과 길이 생략
	c := [5]int{32, 29, 78, 16, 81}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
