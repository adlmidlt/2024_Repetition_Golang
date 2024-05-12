package lesson_6

import "fmt"

func EditMassive() {
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := arr1[:3]
	arr1[1] = "X"
	fmt.Println(arr1)
	fmt.Println(slice1)

	arr2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := arr2[2:5]
	slice2[1] = "X"
	fmt.Println(arr2)
	fmt.Println(slice2)

	arr3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := arr3[:3]  // тут нарезали a, b, c
	slice4 := arr3[2:5] // тут нарезали c, d, e
	arr3[2] = "X"       // поменяли на 18строчке на X -> (a, b, X, d, e) и значение,
	// потом нарезали 19 и 20 строчку, и они пересеклись.
	fmt.Println(arr3)
	fmt.Println(slice3, slice4)
}
