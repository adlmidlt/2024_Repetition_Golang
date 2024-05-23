package lesson_11

import (
	"fmt"
	"headfirstgo/lesson_11/mypkg"
)

func StartExample() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
