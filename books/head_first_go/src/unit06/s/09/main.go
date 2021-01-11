package main

import "fmt"

// 슬라이스를 사용할 때 정말로 주의해야하는 점
// append를 사용할 때는 반드시 append에 전달한 것과 동일한 슬라이스 변수에
// 재할당을 하도록 하자.
func main() {
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}
