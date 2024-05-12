package lesson_5

import (
	"fmt"
	"headfirstgo/lesson_5/datafile"
	"log"
)

func Average() {
	numbers, err := datafile.GetFloats("lesson_5/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
}
