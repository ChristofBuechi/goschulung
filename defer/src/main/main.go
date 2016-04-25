package main

import (
	"fmt"
)


func calculateFunc(a int, b int)  int {
	return a+b
}

func findMax(a []int) (max int) {
	for _, v := range a{
		if v > max {
			max = v
		}
	}
	return
}


func main() {

	//named return value: max
	max := findMax([] int {1,2,3,4,99,5})
	fmt.Printf("Biggest value:  %d\n", max)


	//defer example - defer is called late, but with the value in order
	value := 3
	defer fmt.Printf("I'm done. Value for c = %d\n", value)
	value = calculateFunc(value, 2)
	fmt.Printf("Return value: %d\n", value)
}

