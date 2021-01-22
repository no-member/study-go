package main

import (
	"log"
	"unit10/calendar"
)

func main() {
	event := calendar.Event{}
	if err := event.SetTitle("Mom's birthday"); err != nil {
		log.Fatal(err)
	}
}
