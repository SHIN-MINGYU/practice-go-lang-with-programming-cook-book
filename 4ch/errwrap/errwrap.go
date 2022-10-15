package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError func show how to wrap an error and denote remark in error

func WrappedError(err error) error {
	return errors.Wrap(err, "An error is occured in WrappedError")
}

// ErrorTyped structure is erorr used in this package
type ErrorTyped struct {
	error
}

//Wrap func is called when an error is wrapping

func Wrap()  {
	err := errors.New("standard error")
	fmt.Println("Regular error - ", WrappedError(err))
	fmt.Println("Typed error - ", WrappedError(ErrorTyped{errors.New("typed error")}))
	fmt.Println("Nil - ", WrappedError(nil))
}


