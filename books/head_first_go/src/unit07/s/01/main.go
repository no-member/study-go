package main

import "fmt"

// make를 사용해서 생성해야하는 map
func main() {
	var ranks map[string]int
	ranks = make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["silver"])
}
