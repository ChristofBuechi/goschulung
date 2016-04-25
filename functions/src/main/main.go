package main

import (
	"fmt"
)

func helper(a int, b string) (int, string)  {
	return a, b
}


func main() {

	val1, val2 := helper(3, "some string")
	fmt.Printf("Number: %d, String: %s\n", val1, val2)


	
}

