package main

import "fmt"

<<<<<<< HEAD
// 슬라이드의 변경
func main() {
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	fmt.Println("array1:", array1)
	fmt.Println("slice1:", slice1)
	fmt.Println()

	array1[0] = "X"
	fmt.Println("array1:", array1)
	fmt.Println("slice1:", slice1)
	fmt.Println()

	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	fmt.Println("array2:", array2)
	fmt.Println("slice2:", slice2)
	fmt.Println()

	slice2[1] = "Y"
	fmt.Println("array2:", array2)
	fmt.Println("slice2:", slice2)
	fmt.Println()

	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3, slice4)
}
