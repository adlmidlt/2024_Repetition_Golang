package lesson_1

import (
	"fmt"
	"reflect"
)

func Transformation() {
	var length = 1.2
	var width = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))

	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(float64(width)))
}
