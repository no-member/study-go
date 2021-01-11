package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {

	var hello func(a int, b int) int = sum
	world := sum

	fmt.Println(hello(2, 3))
	fmt.Println(world(2, 3))
}
