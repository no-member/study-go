package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		goto ERROR
		// start
		b := 1
		fmt.Println(b)
		// end
		// start와 end 사이는 출력 되지 않음
	}

ERROR:
	fmt.Println("Error")
}
