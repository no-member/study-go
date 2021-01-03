package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int
	fmt.Sscanf("57", "%d", &num)

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(num)
	fmt.Println()

	var s string
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	fmt.Println()

	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)
	fmt.Println()
}
