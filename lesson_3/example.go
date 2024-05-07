package lesson_3

import "fmt"

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func Example() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
