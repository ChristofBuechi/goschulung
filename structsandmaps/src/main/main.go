package main

import (
	"fmt"
)

var (
	mySlice[]int
	myMap map[string]int
)

func main() {
	myMap := make(map[string]int)
	myMap["zero"] = 0
	myMap["first"] = 1
	myMap["second"] = 2

	for key, val := range myMap {
		fmt.Printf("%s: %d\n", key, val)
	}

	val := myMap["second"]
	fmt.Printf("The value is %d\n", val)

	unknown, found := myMap["unknwon"]; if found {
		fmt.Printf("The value is %d\n", unknown)
	} else {
		fmt.Println("The value is not found")
	}
}

