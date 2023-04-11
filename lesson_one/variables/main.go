package main

func main() {

	var i int = 10 // зависит от платформы 32 или 64 бита
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1          // int8, int16, int32, int64
	var unsignedInt uint = 100500         // зависит от платформы 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1 // uint8, uint16, uint32, uint64

	// fmt.Printf("i = %v\n autoInt = %v\n bigInt = %v\n unsignedInt = %v\n unsignedBigInt = %v\n", i, autoInt, bigInt, unsignedInt, unsignedBigInt)
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)
	// return

	//числа с плавующей точкой
	var p float32 = 3.14 //float = float32, float64
	println("float", p)
	// return

	//булевы переменные
	var b = true
	println("bool variable", b)
	// return

	//строки
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)
	// return

	//бинарные данные
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)
	// return

	/*
		короткое объявление
	*/
	meaningOfLife := 42 // для создания только новой переменной
	println("Meaning of life is ", meaningOfLife)
	// return

	/*
		 приведение типов
	*/
	println("float to int conversion ", int(p))
	println("int to string conversion ", string(48))
	// return

	/*
		комплексные числа
	*/
	z := 2 + 3i
	println("complex number: ", z)
	// return

	/*
		операции со строками
	*/
	s1 := "Vasily"
	s2 := "Romanov"
	fullName := s1 + " " + s2
	println("name length is:", fullName, len(fullName))

	escaping := `Hello\r\n 
	World`
	println("as-is escaping: ", escaping)
	// return

	/*
		Значения по умолчанию
	*/

	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("dafault values: ", defaultInt, defaultFloat, defaultString, defaultBool)
	// return

	/*
		несколько переменных
	*/

	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)
	println(m0, m2, m3)
	return

}
