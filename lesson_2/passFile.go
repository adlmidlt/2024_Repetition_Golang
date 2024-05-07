package lesson_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// PassFile - сообщает, сдал ли пользователь экзамен.
func PassFile() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
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
