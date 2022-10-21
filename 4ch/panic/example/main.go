package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/4ch/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}