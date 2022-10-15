package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/4ch/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()

}