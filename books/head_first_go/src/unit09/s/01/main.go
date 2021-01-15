package main

import (
	"fmt"
)

type Literts float64
type Gallons float64

func main() {
	// var carFuel Gallons
	// var busFuel Literts
	//
	// carFuel = Gallons(10.0)
	// busFuel = Literts(240.0)

	// 단축 선언
	carFuel := Gallons(10.0)
	busFuel := Literts(240.0)
	fmt.Println(carFuel, busFuel)
	fmt.Println()

	carFuel = Gallons(Literts(40.0) * 0.264)
	busFuel = Literts(Gallons(63.0))
	fmt.Printf("Gallons: %0.1f Literts: %0.1f\n", carFuel, busFuel)
}
