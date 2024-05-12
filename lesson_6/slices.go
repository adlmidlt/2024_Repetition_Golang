package lesson_6

import "fmt"

func Slices() {
	sum := 0
	mySlice := []int{731, 900, 450, 851, 731, 1181, 748, 450}
	for sl := range mySlice {
		sum += mySlice[sl]
	}
	fmt.Println(sum, len(mySlice))
}

// BaseMassive 217 страница
func BaseMassive() {
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	fmt.Println(slice1)
}
