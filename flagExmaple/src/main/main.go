package main

import (
	"fmt"
	"flag"
	"log"
)

var (
	myname = flag.String("my_name", "", "the name of the user")
)

func main() {
	flag.Parse()

	if *myname == "" {
		log.Fatal("Hey.. you didn't tell me your name")
	}
	fmt.Printf("name has value %s!\n", *myname)

}


