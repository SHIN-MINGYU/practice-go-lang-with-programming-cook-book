package main

import "github.com/go-practice-width-cook-book/4ch/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}