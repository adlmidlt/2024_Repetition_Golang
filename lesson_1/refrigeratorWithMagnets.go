package lesson_1

import "fmt"

// RefrigeratorWithMagnets - холодильник с магнитами.
func RefrigeratorWithMagnets() {
	var originalCount = 10
	var eatenCount = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
