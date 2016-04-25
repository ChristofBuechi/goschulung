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

	var a []string

//	for key, val := range myMap {
//		a[val] = key
//	}

	for key := range myMap {
		a = append(a,key)
	}



	for i:= 0; i < len(a); i++ {
		fmt.Printf("Index: %d, Value %s\n", i, a[i])
	}


}

