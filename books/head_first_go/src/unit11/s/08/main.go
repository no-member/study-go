package main

import "unit11/s/01/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	// 런타임 패닉 발생!
	TryOut(gadget.TapePlayer{})
}
