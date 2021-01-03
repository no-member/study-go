package main

import "fmt"

// printf의 사용법
func main() {
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Println()

	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Println()

	fmt.Printf("That will be $%f please.\n", 0.23*5)
}
