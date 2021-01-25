package main

import (
	"fmt"
	"unit11/s/01/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Play")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	TryOut(gadget.TapeRecorder{})
	fmt.Println()

	TryOut(gadget.TapePlayer{})
}
