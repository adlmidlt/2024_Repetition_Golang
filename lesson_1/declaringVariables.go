package lesson_1

import (
	"fmt"
	"reflect"
)

// DeclaringVariablesFirstWay - 1 способ объявления переменных, если мы не знаем значение
func DeclaringVariablesFirstWay() {
	var quantity int          // по умолчанию 0
	var length, width float64 // по умолчанию 0
	var customerName string   // по умолчанию "пустая строка"

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}

// DeclaringVariablesSecondWay 2 способ объявление переменных, если мы знаем занчение
func DeclaringVariablesSecondWay() {
	var quantity = 4
	var length, width = 1.2, 2.4
	var customerName = "Damon Cole"
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))
}

// DeclaringVariablesThirdWay - 3 способ объявления переменных, короткий.
func DeclaringVariablesThirdWay() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
