package main

import ("fmt"
	"io/ioutil"
	"strings"
	"os"
)

var (
	line string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	fmt.Println("read from file")

	byteArray, err := ioutil.ReadFile("/etc/fstab");
	check(err)
	line = string(byteArray[:])

	words := strings.Split(line, " ")
	for _,w := range words {
		fmt.Printf("%s \n", w)
	}

	fmt.Println(line)


	fmt.Println("\nos open\n")
	f, err := os.Open("/etc/fstab")
	check(err)

	b1 := make([]byte, 1000)
	n1, err := f.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}