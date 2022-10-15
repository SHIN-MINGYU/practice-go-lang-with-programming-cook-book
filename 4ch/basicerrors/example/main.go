package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/4ch/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error", err)
}