package lesson_3

import "fmt"

func CopyArguments() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
	fmt.Println(&amount)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	fmt.Println(&myBool)

}

func double(number *int) {
	*number *= 2
}

func createPointer() *float64 {
	myFloat := 98.5

	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}
