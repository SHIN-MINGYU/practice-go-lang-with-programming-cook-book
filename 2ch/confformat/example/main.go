package main

import "github.com/go-practice-width-cook-book/2ch/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.UnmarshalAll(); err!= nil {
		panic(err)
	}
	if err:= confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}