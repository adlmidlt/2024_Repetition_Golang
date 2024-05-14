package lesson_7

import (
	"fmt"
	"headfirstgo/lesson_5/datafile"
	"log"
)

func Count() {
	lines, err := datafile.GetStrings("lesson_7/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}

// CountVoicesHard 247 страница
func CountVoicesHard() {
	lines, err := datafile.GetStrings("lesson_7/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

func CountVoicesEasy() {
	lines, err := datafile.GetStrings("lesson_7/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for k, v := range counts {
		fmt.Printf("Voites for %s: %d\n", k, v)
	}

}

func Maps() {
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap["a"])

}

func Status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade found for %s\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is falling!\n", name)
	}
}

func Example() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, line := range data {
		counts[line]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("No count for letter %s\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
	delete(counts, "a")
	fmt.Println(counts)
}
