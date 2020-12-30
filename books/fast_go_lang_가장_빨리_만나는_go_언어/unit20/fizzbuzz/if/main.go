package main

import "fmt"

func main() {
	// if를 사용
	for i := 0; i < 100; i++ {
		if (i+1)%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if (i+1)%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		if (i+1)%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		fmt.Println(i + 1)
	}
}
