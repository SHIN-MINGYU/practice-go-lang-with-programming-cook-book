package main

import "github.com/go-practice-width-cook-book/1ch/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}