package dataconv

import "fmt"

// CheckType function print based on interface type
func CheckType(v interface{}){
	switch v.(type) {
    case int:
        fmt.Println("this is int!")
	case string:
		fmt.Println("this is string!")
	default:
		fmt.Println("this is not expected type!")
	}
}


// Interfaces function is show how to convert type to another type from annonymous interface 

func Interfaces() {
    CheckType("test")
	CheckType(1)
	CheckType(true)
    CheckType(3.14)

	var i interface{}
	i = "test"

	// check interface manually
	if val, ok := i.(string); ok {
		fmt.Println("val is ", val)
	}

	// this code should failed 
	if _ ,ok := i.(bool); !ok {
		fmt.Println("this is not expected type!")
	}
}