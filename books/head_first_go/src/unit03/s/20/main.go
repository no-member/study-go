package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println()

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	fmt.Println()

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
	fmt.Println()

	myNewInt := 10
	fmt.Println(myNewInt)
	myNewIntPointer := &myNewInt
	*myNewIntPointer = 8
	fmt.Println(*myNewIntPointer)
	fmt.Println(myNewInt)
}
