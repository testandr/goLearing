package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	var res float64
	if x < 0 {
		return math.NaN()
	} else if x == 0 {
		return x
	} else {
		zeroPnt := x / 2
		for i := 0; i <= 6; i++ {
			interRes := (zeroPnt + x/zeroPnt) / 2
			zeroPnt = interRes
			res = interRes
		}
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
}
