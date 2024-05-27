package example

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Horn!")
}

type NoiseMarker interface {
	MakeSound()
}

type Anything interface {
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

type Robot string

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func Play(n NoiseMarker) {
	n.MakeSound()

}
