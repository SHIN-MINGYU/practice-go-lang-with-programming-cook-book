package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listening on port : 8000")
	// For dealing with all request, mount one handler on localhost:8000
	log.Panic(http.ListenAndServe("localhost:8000",http.HandlerFunc(wsHandler)))
}