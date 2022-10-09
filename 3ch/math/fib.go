package math

import (
	"math/big"
)

// global variables for saving fibonacci numbers
var memorize map[int]*big.Int

func init() {
// initialize memoization
	memorize = make(map[int]*big.Int)
}


//Fib functionis print n'th number in fibonacci notation
// if the number is less than zero, return 1

// this funtion calculate recusively, use big.Int because of int64's stack overflow 
func Fib(n int) *big.Int {
	if n < 0 {
        return big.NewInt(1)
    }

	// basic case
	if n < 2 {
		memorize[n] = big.NewInt(1)
	}

	//  check save before, 
	// if had saved value, not calculate, return saved value
	if val, ok := memorize[n]; ok {
		return val
	}

	// initailize map, add two fibonacci value before
	memorize[n] = big.NewInt(1)
    memorize[n].Add(memorize[n], Fib(n-1))
	memorize[n].Add(memorize[n], Fib(n-2))
	return memorize[n]
}