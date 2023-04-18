package main

import "fmt"

func main() {
	var mm map[string]string
	fmt.Println("uninitialized map", mm)
	// при попытке добавить в мапу элемент будет вызвана паника
	// nm["test"] = "ok"

	// полная инициализация
	// var mm2 map[string]string = map[string]string{}
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// короткая инициализация
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Vasily"
	fmt.Println(mm3)

	//получение значения
	firstName := mm3["firstName"]
	fmt.Println("firstName", firstName, len(firstName))

	// что будет если обратиться ко значению которого нет
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	// проверка на то, что есть ли значение
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist:", ok)

	// получение только признака существования
	_, exist := mm3["firstName"]
	fmt.Println("firstName exist", exist)

	// удаление значения
	delete(mm3, "firstName")
	_, exist = mm3["firstName"]
	fmt.Println("firstName exist:", exist)

	// так же можно передать переменную с ключем
	/*
		key := "firstName"
		delete(mm3, key)
	*/

	// копировать мапу нельзя, только через цикл...
	// порядок в мапе неопределен...

}
