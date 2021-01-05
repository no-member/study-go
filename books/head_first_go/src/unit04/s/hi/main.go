package main

import (
	"fmt"
	"unit04/s/greeting"
	"unit04/s/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}
