package lesson_2

import (
	"fmt"
	"headfirstgo/lesson_4/keyboard"
	"log"
)

// PassFile - сообщает, сдал ли пользователь экзамен.
func PassFile() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing!"
	} else {
		status = "falling!"
	}

	fmt.Println("A grade of", grade, "is", status)
}
