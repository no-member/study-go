package main

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday
		Tuesdday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Sunday)
	fmt.Println(Friday)
	fmt.Println()

	const (
		a = iota * 30
		b = iota * 30
		c = iota * 30
		d = iota * 30
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
