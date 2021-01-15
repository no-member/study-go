package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

// 외부 구조체 임베딩
func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	// 임베딩된 구조체를 사용하기 subscriber.Address.City의 형태로 접근할 필요가
	// 없이 subscriber.City의 형태로 접근 가능
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "618111"
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)
}
