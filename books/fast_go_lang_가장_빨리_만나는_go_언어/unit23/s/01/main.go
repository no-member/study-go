package main

import "fmt"

func main() {
	var a map[string]int = make(map[string]int)
	var b = make(map[string]int)
	c := make(map[string]int)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
