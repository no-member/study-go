package main

import "fmt"

// map 리터럴 사용하기
func main() {
	ranks := map[string]int{
		"bronze": 3,
		"silver": 2,
		"gold":   1,
	}

	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])
	fmt.Println()

	elements := map[string]string{"H": "Hydrogen", "Li": "Lithium"}
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
}
