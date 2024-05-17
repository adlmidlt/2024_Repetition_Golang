package lesson_8

import "fmt"

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name   string
	Salary float64
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func StartEmployee() {
	address := Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	subscriber := Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber)

	//var employee Employee
	//employee.Name = "Joy Carr"
	//employee.Salary = 60000
	//fmt.Println(employee.Name)
	//fmt.Println(employee.Salary)
}
