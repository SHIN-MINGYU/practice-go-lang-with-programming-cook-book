package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap func
// unlock error's wrapping
// excute erorr's type assertion

func Unwrap()  {
	err := error(ErrorTyped{errors.New("an error is occured")})
	err = errors.Wrap(err,"wrapped")
	fmt.Println("wrapped error : ", err)

	// can handle various types of error
	switch errors.Cause(err).(type){
	case ErrorTyped:
		fmt.Println("typed error : ", errors.Cause(err))
	default:
		fmt.Println("unknown error : ", errors.Cause(err))
	}
}

// StackTrace func
// @return all stack about errors

func StackTrace() {
	err := error(ErrorTyped{errors.New("an error is occured")})
	err = errors.Wrap(err,"wrapped")

	fmt.Printf("%+v\n",err)
}
