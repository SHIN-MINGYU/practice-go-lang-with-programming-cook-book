package main

import (
	"flag"
	"fmt"
)

func main() {
	// initailize setting
	c := Config{}
	c.Setup()

	// normally call in main function
	flag.Parse()

	fmt.Println(c.GetMessage())

}
