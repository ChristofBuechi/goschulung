package main

import (
	"fmt"
)


var (
	myArray = []int{1,2,3,4}
)

func main() {

	for i:= 0; i < len(myArray); i++ {
		fmt.Printf("Index: %d, Value %d\n", i, myArray[i])
	}



}