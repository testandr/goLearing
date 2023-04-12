package main

func main() {

	/*
		+++++++++++++ УСЛОВИЯ
	*/

	a := true
	if a {
		println("hello world")
	}

	b := 1
	c := 22
	if (b == 1 && a) || c != 22 {
		println("неявное преобразование (if b) не работает")
	}

	mm := map[string]string{
		"firstName": "Vasily",
		"lastName":  "Romanov",
	}
	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exist, = ", firstName)
	} else {
		println("no firstName")
	}

	// другой вариант записи if условия из строки 20
	// if firstName, ok := mm["firstName"]; firstName != "Vasily" {
	// 	println("firstName key exist, = ", firstName, ok)
	// } else {
	// 	println("no firstName")
	// }

	// другой вариант записи if условия из строки 20
	// if _, ok := mm["firstName"]; ok" {
	// 	println("firstName key exist, = ", ok)
	// } else {
	// 	println("no firstName")
	// }

	if firstName, ok := mm["firstName"]; !ok {
		println("no firstName")
	} else if firstName == "Vasily" {
		println("firstName is Vasily")
	} else {
		println("firstName is not Vasily")
	}

	/*
		+++++++++++++ ЦИКЛЫ
	*/

	for {
		println("бесконечный цикл")
		break
	}

	sl := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0

	// операции по slice

	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl[idx]
		idx++
		println("while-stype loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		println("range slice by index", idx)
	}

	for idx, val := range sl {
		println("range slice by idx-value", idx, val)
	}

	for _, val := range sl {
		println("range slice by idx-value", val)
	}

	for idx, _ := range sl {
		println("range slice by idx-value", idx)
	}

	// операции по maps

	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key-value", key, val)
	}

	for _, val := range mm {
		println("range map by value", val)
	}

	for key, _ := range mm {
		println("range map by key", key)
	}

	/*
		+++++++++++++ SWITCH
	*/

	mm["firstName"] = "Vasily"
	mm["flag"] = "Ok"
	switch mm["firstName"] {
	case "Vasily", "Evgeny":
		println("switch - name is Vasily")
		// в отличии от других языков - не переходим в другой вариант по-умолчанию
	case "Petr":
		if mm["flag"] == "Ok" {
			break // выходим из switch, чтобы не выполнять переход в другой вариант
		}
		println("switch - name is Petr")
		fallthrough // переходим в следующий вариант
	default:
		println("switch - some another name")
	}

	// как замена множественным if else
	switch {
	case mm["firstName"] == "Vasily":
		println("switch2 - Vasily")
	case mm["lastName"] == "Romanov":
		println("switch2 - Romanov")
	}

	// выход из цикла когда switch внутри него делается при помощи лейбла
MyLoop:
	for key, val := range mm {
		println("switch in loop", key, val)
		switch {
		case key == "firstName" && val == "Vasily":
			println("switch - break loop here")
			break MyLoop
		}
	}
}
