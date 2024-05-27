package lesson_11

import "headfirstgo/lesson_11/example"

func StartMakeSound() {
	var toy example.NoiseMarker
	toy = example.Whistle("Toyco Canary")
	toy.MakeSound()

	toy = example.Horn("Toyco Blaster")
	toy.MakeSound()

	example.Play(example.Whistle("Toyco Canary"))
	example.Play(example.Horn("Toyco Blaster"))
}
