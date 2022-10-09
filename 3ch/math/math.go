package math

import (
	"fmt"
	"math"
)

// Examples function show menual of some math package's functions
func Examples() {
	// sqrt 
	i := 25
	// i is int, so convert
	result := math.Sqrt(float64(i))

	// 25's square root is 5
	fmt.Println(result)
	
	// Ceil calculate
	result = math.Ceil(9.3)
	
	fmt.Println(result)

	// Floor calculate
    result = math.Floor(9.3)
	fmt.Println(result)

	// also math package offer some const
	const pi = math.Pi
    fmt.Println(pi)
	const e = math.E
	fmt.Println(e)
}