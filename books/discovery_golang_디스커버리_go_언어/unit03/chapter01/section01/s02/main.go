package main

import "fmt"

func main() {
	for _, r := range "가힣" {
		fmt.Println(string(r), r)
	}
	fmt.Println()
}
