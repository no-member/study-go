package main

import "fmt"

// 키와 값에 여러 타입을 가질 수 있는 맵
func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
	fmt.Println()

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
}
