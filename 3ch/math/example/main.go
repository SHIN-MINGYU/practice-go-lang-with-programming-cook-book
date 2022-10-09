package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/3ch/math"
)

func main() {
	math.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}