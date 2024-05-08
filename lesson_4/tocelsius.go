package lesson_4

import (
	"fmt"
	"headfirstgo/lesson_4/keyboard"
	"log"
)

func ToCelsius() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees celsius", celsius)
}
