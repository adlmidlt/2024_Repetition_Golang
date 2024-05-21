package lesson_10

import (
	"fmt"
	"headfirstgo/lesson_10/calendar"
	"log"
)

func StartDate() {
	date := calendar.Date{}
	var err error
	err = date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(22)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}
