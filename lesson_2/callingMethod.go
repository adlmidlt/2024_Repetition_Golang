package lesson_2

import (
	"fmt"
	"strings"
	"time"
)

func CallingMethod() {
	var now = time.Now()
	var year = now.Year()
	fmt.Println(year)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
