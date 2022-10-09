package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/2ch/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text : "I'm Red!",
	}

	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "I'm Green!"

	fmt.Println(r.String())

	r.TextColor = ansicolor.Black
	r.Text = "I'm None!"

	fmt.Println(r.String())

}