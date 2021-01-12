package main

import "fmt"

func main() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 89.99

	clothing := map[string]float64{"pants": 89.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["Necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}
