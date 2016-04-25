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
		log.Fatal("Hey.. you didn't tell me your name\n")
	}

	if len(*myname) < 2 {
		log.Fatal("I dont trust you\n")
	}

	switch *myname {
	case "christof":
		fmt.Print("welcome back christof\n")
	case "unknown":
		fmt.Print("you are unknown - really ?\n")
	}
	fmt.Printf("name has value %s!\n", *myname)

}


