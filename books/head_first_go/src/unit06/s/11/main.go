package main

import "fmt"

func main() {
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println()

	fmt.Println(len(intSlice))
	fmt.Println()

	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v", intSlice)
	fmt.Println()
}
