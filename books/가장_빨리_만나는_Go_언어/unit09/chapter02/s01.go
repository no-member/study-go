package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s3 + s2)
	fmt.Println("안녕하세용" + s1)
}
