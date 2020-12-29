package main

import "fmt"

func main() {
	singTazan(4)
}

func singTazan(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("타잔이 ", 10*(i+1), "원짜리 팬티를 입고,")
		fmt.Println(10*(i+2), "원짜리 칼을 차고 노래를 한다.")
	}
}
