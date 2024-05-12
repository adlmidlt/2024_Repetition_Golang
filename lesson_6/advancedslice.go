package lesson_6

import "fmt"

// AdvancedSlice 219 страница
func AdvancedSlice() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))
}
