package main

import "fmt"

func main() {
	i := 10

	if i > 5 {
		fmt.Println("5 초과")
	} else if i == 5 {
		fmt.Println("5")
	} else {
		fmt.Println("5 미만")
	}
}
