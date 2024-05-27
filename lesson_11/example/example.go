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
type Robot string

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func Play(n NoiseMarker) {
	n.MakeSound()

}
