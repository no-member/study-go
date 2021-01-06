package main

import "fmt"

// 배열을 만들 때 배열에 포함된 모든 값은 배열이 가진 탕비의 제로 값으로 초기화
// 됩니다.
func main() {
	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])
	fmt.Println()

	var notes [7]string
	notes[0] = "do"
	fmt.Println(notes[3])
	fmt.Println(notes[6])
	fmt.Println(notes[0])
	fmt.Println()

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
}
