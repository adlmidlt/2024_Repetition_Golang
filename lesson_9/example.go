package lesson_9

import "fmt"

type Number int

func Example() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}

func (n *Number) Add(otherNumber int) {
	fmt.Println(*n, "plus", otherNumber, "is", int(*n)+otherNumber)
}

func (n *Number) Subtract(otherNumber int) {
	fmt.Println(*n, "plus", otherNumber, "is", int(*n)-otherNumber)
}
