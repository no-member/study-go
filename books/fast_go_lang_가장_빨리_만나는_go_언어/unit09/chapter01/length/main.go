package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println("len(s1) : ", len(s1))                                  // 6 <= len의 경우 글자의 (바이트의 길이)를 구한다.
	fmt.Println("len(s2) : ", len(s2))                                  // 5
	fmt.Print("utf8.RuneCountInString(s1)", utf8.RuneCountInString(s1)) // 2 <= 문자열의 실제 길이를 구함
}
