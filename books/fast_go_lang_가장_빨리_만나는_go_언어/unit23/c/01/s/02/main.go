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

	// 맵에 데이터가 있는 검사할 때는 두 번째 변수에 키의 존재 여부가 저장된다.
	value, ok := solarSystem["Pluto"]
	fmt.Println(value, ok)

	// 값이 존재하는지 확인한 뒤에 출력하기
	if value, ok := solarSystem["Saturn"]; ok {
		fmt.Println(value)
	}

	if value, ok := solarSystem["Pluto"]; ok {
		fmt.Println(value)
	}
}
