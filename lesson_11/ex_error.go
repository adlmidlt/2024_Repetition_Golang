package lesson_11

import (
	"fmt"
	err2 "headfirstgo/lesson_11/err"
	"log"
)

func StartExampleError() {
	var err error
	err = err2.ComedyError("what's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	err = err2.CheckTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
