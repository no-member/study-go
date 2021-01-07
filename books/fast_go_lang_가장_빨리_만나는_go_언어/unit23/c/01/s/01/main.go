package main

import "fmt"

// 맵에 데이터 저장하기
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

	fmt.Println(solarSystem)
	fmt.Println(solarSystem["Earth"])
	fmt.Println(solarSystem["Pluto"]) // 0을 출력 => 값의 제로 값이 출력됨
}
