package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println("пустой слайс", s1)
	s1 = append(s1, 100) // добавление элемента в слайс
	fmt.Println("уже не пустой слайс", s1)
	fmt.Println("длина слайса", len(s1))
	fmt.Println("Длина внутреннего массива в слайсе", cap(s1))

	s1 = append(s1, 102)
	fmt.Println("Длина внутреннего массива в слайсе", s1, cap(s1))
	s1 = append(s1, 103)
	s1 = append(s1, 104)
	fmt.Println("Длина внутреннего массива в слайсе", s1, cap(s1))
	s1 = append(s1, 105)
	fmt.Println("Длина внутреннего массива в слайсе", s1, cap(s1))

	s12 := []int{10, 20, 30}
	fmt.Println(s12)

	s13 := []int{
		11,
		21,
		31,
	}
	fmt.Println(s13)

	// добавить слайс в слайс
	s1 = append(s1, s12...) //так как ждет тип int
	fmt.Println(s1)

	var slm [][]int
	slm = append(slm, s12) // так как ждет тип слайс интов "[]int"
	fmt.Println(slm)

	// создать слайс нужной длины
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// создать слайс с нужной длиной и размером
	slice4 := make([]int, 10, 15)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))

	//внутри слайса - ссылка на массив, она копируеться если просто присвоить
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)

	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	slice4[1] = 999
	fmt.Println(slice4, slice5)

	//правильное копирование слайса
	//необходимо создавать для копирования такой же массив как и тот который будет скопирован
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	//вырезка определенных кусков

	fmt.Println("Вырезка куска с 1 по 4 элемент включительно из слайса slice7", slice7[1:5])
	fmt.Println("Вырезка куска не указывая откуда по 1 элемент включительно из слайса slice7", slice7[:2])
	fmt.Println("Вырезка куска с 10 и до конца из слайса slice 7", slice7[10:])

	//создание слайса из массива
	a := [...]int{5, 6, 7}
	s18 := a[:]
	a[1] = 8
	fmt.Println("слайс из массива", s18)
}
