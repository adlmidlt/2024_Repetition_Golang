package lesson_3

import "fmt"

func RefrigeratorWithMagnets() {
	myInt := 42
	myIntPtr := &myInt
	fmt.Println(myIntPtr)
	fmt.Println(*myIntPtr)
}
