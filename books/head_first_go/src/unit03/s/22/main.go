package main

import "fmt"

func main() {
	myBool := true
	printPointer(&myBool)
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}
