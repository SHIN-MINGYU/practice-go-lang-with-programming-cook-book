package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/4ch/log"
)

func main() {
	fmt.Println("basic logging and modification of logger : ")
	log.Log()
	fmt.Println("logging handled errors : ")
	log.FinalDestination()
}