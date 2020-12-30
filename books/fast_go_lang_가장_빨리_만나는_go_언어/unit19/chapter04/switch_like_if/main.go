package main

import "fmt"

func main() {
	i := 7

	switch {
	case 5 <= i && i < 10:
		fmt.Println("5 이상, 10 미만")
	case 0 <= i && i < 5:
		fmt.Println("0 이상, 10 미만")
	}
}
