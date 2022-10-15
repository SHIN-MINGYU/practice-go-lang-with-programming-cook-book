package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue is the way to create package level erorr for checking
// so if error == ErrorValue
var ErrorValue = errors.New("this is typed error")

// TypedError is the way to create error type for expression(err.(type) === ErrorValue)
type TypedError struct{
	error
}

// BasicErrors function show how to create some erorrs
func BasicErrors(){
	err := errors.New("this is quick and easy way to create an error")
	fmt.Println("errors.New : ",err)

	err = fmt.Errorf("an error occurred : %s", "somthing")
	fmt.Println("fmt.Errorf : ",err)

	err = ErrorValue
	fmt.Println("ErrorValue : ",err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("TypedError : ",err)
}