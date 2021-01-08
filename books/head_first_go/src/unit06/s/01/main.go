package main

import "fmt"

func main() {
	var note []string
	note = make([]string, 7)

	note[0] = "do"
	note[1] = "re"
	note[2] = "mi"
	fmt.Println(note[0])
	fmt.Println(note[1])
}
