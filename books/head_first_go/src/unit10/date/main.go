package main

import (
	"fmt"
	"log"
	"unit10/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
