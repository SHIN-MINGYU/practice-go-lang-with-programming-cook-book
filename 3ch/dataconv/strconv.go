package dataconv

import (
	"fmt"
	"strconv"
)

// Strconv function show example of some strconv function
func Strconv() error{
	// strconv is good way to convert to string from another type or to  another type from string
	s := "1234"

	// set 10 in demical setting, can specify 64bit in accurency setting 
	res, err := strconv.ParseInt(s, 10, 64)
	if err!= nil {
        return err
    }
	
	fmt.Println(res)

	// try hexadecimal conversion
	res, err = strconv.ParseInt(s, 16, 64)
    if err!= nil {
        return err
    }

	fmt.Println(res)

	// also can convert to other types
	val, err := strconv.ParseBool("true")
	if err!= nil {
        return err
    }

	fmt.Println(val)

	return nil
}