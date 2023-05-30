package main

import (
	"fmt"
)

func showMeTheType(i interface{}) string {
	switch t := i.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	case int32:
		return "int32"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case uint:
		return "uint"
	case int8:
		return "int8"
	case string:
		return "string"
	case []int:
		return "[]int"
	case map[string]bool:
		return "map[string]bool"
	default:
		return fmt.Sprintf("Not known type %v", t)
	}

}

func main() {
	// showMeTheType()
}
