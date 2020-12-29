package main

import "fmt"

func main() {
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(6))
}

func fibo(n int) int {
	p, q := 0, 1

	for i := 0; i < n; i++ {
		p, q = q, p+q
	}

	return p
}
