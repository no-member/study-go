package main

import "fmt"

// map 안에 map 만들기
func main() {
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": {
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},
	}

	fmt.Println(terrestrialPlanet)
	fmt.Println(terrestrialPlanet["Mercury"])
}
