package main

import (
	"fmt"
)

func helper(a int, b string) (int, string)  {
	return a, b
}

func printFunction(a int, b int, f func(int, int) int)  {
	v := f(a , b)
	fmt.Printf("value: %d\n", v);
}

func calculateFunc(a int, b int)  int {
	return a+b
}


func main() {

	val1, val2 := helper(3, "some string")
	fmt.Printf("Number: %d, String: %s\n", val1, val2)


	value := calculateFunc(1, 2)
	fmt.Printf("Return value: %d\n", value)

	printFunction(1, 5, calculateFunc);
	printFunction(1, 5, func (a, b int) int {
		return a + b
	})


	
}

