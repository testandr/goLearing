package main

import "fmt"

const size uint = 3

func main() {
	var a1 [3]int
	fmt.Println("массив", a1, "длина", len(a1))

	var a2 [2 * size]bool
	fmt.Println(a2, "длина", len(a2))

	a3 := [...]int{1, 2, 3}
	fmt.Println("Длина при инициализации", a3, "длина", len(a3))

	a3[1] = 12
	fmt.Println("После изменения", a3)

	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("массив массивов", aa)
}
