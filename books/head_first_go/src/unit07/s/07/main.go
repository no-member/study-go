package main

import "fmt"

// map에서 제로 값 확인하는 방법
// map의 키에 접근할 때 두번째의 반환 값이 bool을 반환한다.
func main() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)

	value, ok = counters["b"]
	fmt.Println(value, ok)

	value, ok = counters["c"]
	fmt.Println(value, ok)
}
