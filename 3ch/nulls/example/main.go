package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/3ch/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.PointerEncoding(); err!= nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.NullEncoding(); err!= nil {
	}


}