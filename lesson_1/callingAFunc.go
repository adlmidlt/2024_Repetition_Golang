package lesson_1

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
)

func CallingFunctions() {
	// пробелы ставятся автоматически
	fmt.Println("First arg", "Second arg")

	fmt.Println(math.Floor(2.75))

	// string.Title() -> устаревшая функция, надо использовать cases.Title()
	fmt.Println(cases.Title(language.English).String("head first go"))
}
