package lesson_3

import (
	"fmt"
	"log"
)

func RepeatCode() {
	amount, err := paintNeeded(-4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f litters needed\n", amount)
	}

	fmt.Printf("%12s  |  %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------------")

	fmt.Printf("%12s  |  %2d\n", "Stamps", 50)
	fmt.Printf("%12s  |  %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s  |  %2d\n", "Tape", 99)
}

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width cannot be negative %.2f", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height cannot be negative %2.f", height)
	}
	area := width * height

	return area / 10.0, nil
}
