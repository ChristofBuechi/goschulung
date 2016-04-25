package main

import ("fmt"
	"io/ioutil"
)

var (
	byteArray[]byte
	line string
)



func main() {
	fmt.Println("read from file")

	byteArray, _ = ioutil.ReadFile("/etc/fstab");
	line = string(byteArray[:])

	fmt.Println(line)



}