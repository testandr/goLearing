package main

import (
	"log"
)

type memoiz func(int) interface{}

func main() {

	fibMem := memoiz(fib)
	println(fibMem(7))
}

func memorized(fn memoiz) memoiz {
	cache := make(map[int]interface{})

	return func(input int) interface{} {
		if val, found := cache[input]; found {
			log.Println("Read from cache")
			return val
		}

		result := fn(input)
		cache[input] = result
		return result
	}
}

func fib(num int) int {
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
	return result[len(result)-1]
}
