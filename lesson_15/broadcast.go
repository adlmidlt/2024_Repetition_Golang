package lesson_15

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func StartBroadcast() {
	twice(sayHi)
	twice(sayBye)
}
