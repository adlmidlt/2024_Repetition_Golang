package lesson_4

import (
	"fmt"
	"headfirstgo/lesson_4/dates"
)

func Constants() {
	days := 3
	fmt.Println("You appointment is in", days, "days.")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
