package lesson_8

import "fmt"

var myStruct struct {
	number float64
	word   string
	toggle bool
}

type part struct {
	descriptions string
	count        int
}

func showInfo(p part) {
	fmt.Println("Description:", p.descriptions)
	fmt.Println("Count:", p.count)
}

func StartShowInfo() {
	var bolts part
	bolts.descriptions = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)
}

func minimumOrder(description string) part {
	var p part
	p.descriptions = description
	p.count = 100

	return p
}

func StartMinimumOrder() {
	p := minimumOrder("Hex bolts")
	fmt.Println(p.descriptions, p.count)
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
