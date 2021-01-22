package main

import (
	"fmt"
	"log"
	"unit10/p/02/geo"
)

func main() {
	coordinates := geo.Coordinates{}
	err :=  coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Longitude())
	fmt.Println(coordinates.Latitude())
}
