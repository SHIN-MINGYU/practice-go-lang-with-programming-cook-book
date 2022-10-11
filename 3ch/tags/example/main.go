package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/3ch/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}