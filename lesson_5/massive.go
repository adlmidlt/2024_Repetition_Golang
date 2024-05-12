package lesson_5

import (
	"fmt"
)

func Massive() {
	notes := [7]string{"do", "re", "me", "fa", "so", "la", "ti"}
	for _, note := range notes {
		fmt.Println(note)
	}

	numbers := [3]float64{0.1, 0.2, 0.3}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)

}
