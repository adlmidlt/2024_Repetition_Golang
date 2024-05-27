package lesson_11

import (
	"headfirstgo/lesson_11/example"
)

func StartMakeSound() {
	var toy example.NoiseMarker
	toy = example.Whistle("Toyco Canary")
	toy.MakeSound()

	toy = example.Horn("Toyco Blaster")
	toy.MakeSound()

	example.Play(example.Whistle("Toyco Canary"))
	example.Play(example.Horn("Toyco Blaster"))

	var noiseMaker example.NoiseMarker = example.Robot("Borco Ambler")
	var robot = noiseMaker.(example.Robot)
	robot.Walk()
}

func StartAnything() {
	example.AcceptAnything(3.1415)
	example.AcceptAnything("A string")
	example.AcceptAnything(true)
	example.AcceptAnything(example.Whistle("Toyco Canary"))
}
