package main

import (
	"fmt"
)

var (
	mySlice[]int
	myMap map[string]int
)

type keyValPair struct {
	key string
	val int
}

func doubleVal(k keyValPair) keyValPair {
	k.val = k.val * 2
	return k
}

func pointerdoubleVal(k *keyValPair) *keyValPair {
	k.val = k.val * 2
	return k
}

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

	unknown, found := myMap["unknown"]; if found {
		fmt.Printf("The value is %d\n", unknown)
	} else {
		fmt.Println("The value is not found")
	}

	var a []string

	//commented out because of following range operation with apend - otherwise array is to small
/*	for key, val := range myMap {
		a[val] = key
	}*/

	for key := range myMap {
		a = append(a,key)
	}


	for i:= 0; i < len(a); i++ {
		fmt.Printf("Index: %d, Value %s\n", i, a[i])
	}

	//working with struct and copy ?
	kv := keyValPair{
		key: "some key value",
		val: 99,
	}
	dbl := doubleVal(kv)
	fmt.Printf("Key: %s, Val: %d\n", kv.key, kv.val)
	fmt.Printf("Key: %s, Val: %d\n", dbl.key, dbl.val)

	//working with pointer
	kvPointer := &keyValPair{
		key: "some key value",
		val: 99,
	}

	dblPointer := pointerdoubleVal(kvPointer)
	fmt.Printf("Key: %s, Val: %d\n", dblPointer.key, dblPointer.val)


	for i:= 0; i < len(a); i++ {
		kv = keyValPair{a[i], i}
		fmt.Printf("KeyValuePair: Key: %s, Val: %d\n", kv.key, kv.val)
	}

	var vals []keyValPair
	for key, val := range myMap {
		vals = append(vals, keyValPair{key, val})
	}



}

