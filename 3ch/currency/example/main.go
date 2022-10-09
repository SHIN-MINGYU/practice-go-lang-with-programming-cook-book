package main

import (
	"fmt"

	"github.com/go-practice-width-cook-book/3ch/currency"
)

func main() {
	// 15 dollars 93 cents user input
	userInput := "15.93"

	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err!= nil {
        panic(err)
    }

	fmt.Printf("User input is converted to %d pennies\n", pennies)

	// add 15 cent to pennies
	pennies += 15

	dollars := currency.ConvertPenniesToDollarString(pennies)

	fmt.Printf("Added 15 cents, new value is %s\n", dollars)

}