package log

import (
	"bytes"
	"fmt"
	"log"
)

// Log function use logger setting
func Log() {
	// we will record log in bytes.Buffer
	buf := bytes.Buffer{}

	// secondary argument is prefix, last argument is option to set combine to use 'or' operator
	logger := log.New(&buf,"logger : ",log.Lshortfile|log.Ldate)

	logger.Println("this is a log")

	logger.SetPrefix("new logger : ")

	logger.Printf("you can also add args(%v) and use Fatalln to log and crash", true)

	fmt.Println(buf.String())
}