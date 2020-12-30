package main

import "fmt"

func main() {
	bottles := 99
	for {
		if bottles == 1 {
			fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.")
			fmt.Println("Take one down, pass it around, No more bottles of beer on the wall.")
			bottles--

			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
			fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
			break
		}

		fmt.Println(bottles, "bottles of beer on the wall, ", bottles, "bottles of beer.")
		bottles--

		s := "bottles"
		if bottles == 1 {
			s = "bottle"
		}

		fmt.Println("Take one down, pass it around,", bottles, s, " of beer on the wall.")
	}
}
