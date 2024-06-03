package lesson_14

import (
	"fmt"
	"headfirstgo/lesson_14/proce"
)

func AutoTest() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", proce.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", proce.JoinWithCommas(phrases))
}
