package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}

	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)
	fmt.Println()

	slice2 := underlyingArray[1:4]
	fmt.Println(slice2)

	slice3 := underlyingArray[2:5]
	fmt.Println(slice3)
	fmt.Println()
}
