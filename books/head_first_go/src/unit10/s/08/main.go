package main

import (
	"fmt"
	"log"
	"unit10/calendar"
)

func main() {
	event := calendar.Event{}
	if err := event.SetTitle("Mom's birthday"); err != nil {
		log.Fatal(err)
	}

	if err := event.SetYear(2019); err != nil {
		log.Fatal(err)
	}

	if err := event.SetMonth(5); err != nil {
		log.Fatal(err)
	}

	if err := event.SetDay(27); err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
