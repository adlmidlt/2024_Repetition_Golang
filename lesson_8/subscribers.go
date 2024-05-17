package lesson_8

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true

	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.89
}

func StartSubscriber() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}

func PrintSubscriberStruct() {
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true
	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Rate:", subscriber1.rate)
	fmt.Println("Active?:", subscriber1.active)
}
