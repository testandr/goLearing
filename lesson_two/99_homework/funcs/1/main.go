package main

import (
	"fmt"
	"strconv"
)

// type memoizeFunction func(int, ...int) interface{}

// // TODO реализовать
// var fibonacci memoizeFunction
// var romanForDecimal memoizeFunction

func fib(num int) []int {
	var result []int
	var sum int

	for i := 0; i <= num; i++ {
		if i == 0 || i == 1 {
			result = append(result, i)
		} else {
			sum = result[i-1] + result[i-2]
			result = append(result, sum)
		}
	}
	return result
}

func romNum(num int) string {
	var result string

	numStr := strconv.Itoa(num)
	for i := 0; i < len(numStr); i++ {
		switch len(numStr) - i {
		case 4:
			valInt, _ := strconv.Atoi(string(numStr[i]))
			for y := 0; y < valInt; y++ {
				result += giveLetters(1000)
			}
		case 3:
			valStr := string(numStr[i])
			for y := 0; y < 2; y++ {
				valStr += "0"
			}
			valInt, _ := strconv.Atoi(valStr)
			result += giveLetters(valInt)
		case 2:
			valStr := string(numStr[i])
			for y := 0; y < 1; y++ {
				valStr += "0"
			}
			valInt, _ := strconv.Atoi(valStr)
			result += giveLetters(valInt)
		case 1:
			valStr := string(numStr[i])
			valInt, _ := strconv.Atoi(valStr)
			result += giveLetters(valInt)
		}
	}
	return result
}

func giveLetters(num int) string {
	var result string
	switch {
	case num <= 10:
		switch {
		case num >= 1 && num <= 5:
			switch {
			case 5-num == 0:
				result += "V"
			case 5-num == 1:
				result += "IV"
			case 5-num > 1:
				for i := 0; i < num; i++ {
					result += "I"
				}
			}
		case num > 5 && num <= 10:
			switch {
			case 10-num == 0:
				result += "X"
			case 10-num == 1:
				result += "IX"
			case num > 5:
				result = "V"
				for i := 0; i < num-5; i++ {
					result += "I"
				}
			}
		}
	case num <= 100:
		switch {
		case num > 10 && num <= 50:
			switch {
			case 50-num == 0:
				result += "L"
			case 50-num == 10:
				result += "XL"
			case 50-num > 10:
				for i := 0; i < num/10; i++ {
					result += "X"
				}
			}
		case num > 50 && num <= 100:
			switch {
			case 100-num == 0:
				result += "C"
			case 100-num == 10:
				result += "XC"
			case num > 50:
				result = "L"
				for i := 0; i < (num-50)/10; i++ {
					result += "X"
				}
			}
		}
	case num <= 1000:
		switch {
		case num > 100 && num <= 500:
			switch {
			case 500-num == 0:
				result += "D"
			case 500-num == 100:
				result += "CD"
			case 500-num > 100:
				for i := 0; i < num/100; i++ {
					result += "C"
				}
			}
		case num > 500 && num <= 1000:
			switch {
			case 1000-num == 0:
				result += "M"
			case 1000-num == 100:
				result += "CM"
			case num > 50:
				result = "D"
				for i := 0; i < (num-500)/100; i++ {
					result += "C"
				}
			}
		}
	}
	return result
}

// //TODO Write memoization function

// func memoize(function memoizeFunction) memoizeFunction {
// 	return function
// }

// // TODO обернуть функции fibonacci и roman в memoize
// func init() {
// }

func main() {
	// fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romNum(x))
	}

	// for _, x := range []int{1666, 1997, 2008} {
	// 	fmt.Printf("%4d = %s\n", x, romNum(x))
	// }
	fmt.Printf("fibonacci list of 1 %v\n", fib(0))
	// fmt.Printf("roman variant of the num is %v\n", romNum(1666))
}
