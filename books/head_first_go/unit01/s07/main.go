package main

import (
	"fmt"
)

func main() {
	var orinalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", orinalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", orinalCount-eatenCount, "aplles left.")
}
