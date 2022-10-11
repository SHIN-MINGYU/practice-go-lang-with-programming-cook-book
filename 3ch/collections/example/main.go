package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/3ch/collections"
)

func main() {
	ws := []collections.WorkWith{
		collections.WorkWith{"Example", 1},
		collections.WorkWith{"Example 2", 2},
	}

	fmt.Printf("Initial list : %#v\n", ws)

	// first, list's value convert to lowercase
	ws = collections.Map(ws, collections.LowerCaseData)
	fmt.Printf("List converted to lowercase : %#v\n", ws)

	// second, increase all version
	ws = collections.Map(ws, collections.IncrementVersion)
    fmt.Printf("List converted to increase version : %#v\n", ws)

	// last, delete all version less than 3
	ws = collections.Filter(ws, collections.OldVersion(3))
    fmt.Printf("List converted to delete version less than 3 : %#v\n", ws)

}