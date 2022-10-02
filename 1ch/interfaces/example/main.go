package main

import (
	"bytes"
	"fmt"

	"github.com/go-practice-width-cook-book/1ch/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer = ", out.String())
	fmt.Println(("stdout on PipeExample = "))
	if err := interfaces.PipeExample(); err != nil{
		panic(err)
	}

}