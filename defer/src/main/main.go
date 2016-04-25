package main

import (
	"fmt"
)


func calculateFunc(a int, b int)  int {
	return a+b
}


func main() {

	value := 3

	defer fmt.Printf("I'm done. Value for c = %d\n", value)

	value = calculateFunc(value, 2)

	fmt.Printf("Return value: %d\n", value)
}

