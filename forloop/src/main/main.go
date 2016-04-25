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


	for _, value := range myArray {
		fmt.Printf("Value %d\n", value)
	}

	i := 0
	for {
		fmt.Printf("Value: %d\n", myArray[i])
		i++
		if i == len(myArray) {
			break
		}
	}

	k := 0
	for k < len(myArray) {
		fmt.Printf("Value: %d\n", myArray[k])
		k++
	}



}