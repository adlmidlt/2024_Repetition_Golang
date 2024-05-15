package lesson_8

import "fmt"

var myStruct struct {
	number float64
	word   string
	toggle bool
}

var subscriber struct {
	name   string
	rate   float64
	active bool
}

func Print() {

}

func PrintSubscriberStruct() {
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Rate:", subscriber.rate)
	fmt.Println("Active?:", subscriber.active)
}

func PrintMyStruct() {
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
	fmt.Printf("%#v\n", myStruct)
}
