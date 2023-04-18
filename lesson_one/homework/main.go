package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	numInt := 1
	return numInt
}

func ReturnFloat() float32 {
	var numFloat float32 = 1.1
	return numFloat
}

func ReturnIntArray() [3]int {
	numArray := [3]int{1, 3, 4}
	return numArray
}

func ReturnIntSlice() []int {
	numSlice := []int{1, 2, 3}
	return numSlice
}

func IntSliceToString(slInt []int) string {
	var str string
	for _, v := range slInt {
		str += strconv.Itoa(v)
	}
	return str
}

func MergeSlices(slFloat []float32, slInt []int32) []int {
	var slFloatLen = len(slFloat)
	var slComnined = make([]int, slFloatLen+len(slInt))

	for i := range slFloat {
		slComnined[i] = int(slFloat[i])
	}

	for i := range slInt {
		slComnined[i+slFloatLen] = int(slInt[i])
	}

	return slComnined
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	var result = make([]string, len(input))
	var keys []int

	for i := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for i, key := range keys {
		result[i] = input[key]
	}

	return result
}
