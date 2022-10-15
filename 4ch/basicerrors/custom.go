package basicerrors

import "fmt"

// CustomError is stucture for impling Error interface
type CustomError struct {
	Result string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("there was an error, and the result was : %s", e.Result)
}

// SomeFunc return error
func SomeFunc() error {
	c := CustomError{Result: "this"}
	return c
}