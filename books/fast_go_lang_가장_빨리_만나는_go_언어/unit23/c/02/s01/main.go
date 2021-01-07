package main

import "fmt"

func main() {
	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2876
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	for key, value := range solarSystem {
		fmt.Println(key, value)
	}
	fmt.Println()

	for key, _ := range solarSystem {
		fmt.Println(key)
	}
	fmt.Println()

	for _, value := range solarSystem {
		fmt.Println(value)
	}
	fmt.Println()
}
