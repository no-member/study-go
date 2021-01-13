package main

import "fmt"

// map은 정렬되지 않은 컬렉션이기 때문에 for..range 루프는 맵을 무작위 순서로 순회
func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}
}
