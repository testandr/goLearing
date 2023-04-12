package main

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

func IntSliceToString() string {
	str := "1723100500"
	return str
}

func MergeSlices() []int32 {
	slFloat := []float32{1.1, 2.1, 3.1}
	slInt := []int32{4, 5}
	var slComnined = []int32{}
	for _, v := range slFloat {
		slComnined = append(slComnined, int32(v))
	}

	slComnined = append(slComnined, slInt...)
	return slComnined
}

func GetMapValuesSortedByKey() {

}
