package main

import "fmt"

type Literts float64
type Gallons float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Literts(2.5)

	carFuel += Gallons(Literts(40.0) * 0.264)
	busFuel += Literts(Gallons(30.0))

	carFuel += ToGallons(Literts(40.0))
	busFuel += ToLiters(Gallons(30.0))

	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}

func ToGallons(l Literts) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Literts {
	return Literts(g * 3.785)
}
