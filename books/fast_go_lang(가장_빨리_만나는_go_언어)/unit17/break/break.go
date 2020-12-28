package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 4 {
			break
		}

		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()

Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world")
}
