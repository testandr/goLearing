package main

const someInt = 1
const typeInt int32 = 17
const fullName = "Vasily"

const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota
	two
	_    //пустая переменная. пропуск iota
	four // = 4
)

func main() {

	pi := 3.14

	// тип константы может быть определен во время компиляции
	println(pi + someInt)
}
