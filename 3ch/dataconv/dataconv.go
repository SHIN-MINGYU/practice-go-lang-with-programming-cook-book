package dataconv

import "fmt"

// ShowConv fucntion show conversion example of some types
func ShowConv(){
	// int
	var a = 24

	// float64
	var b = 2.0

	// convert int to float64 for calculate
	c := float64(a) * b
	fmt.Println(c)
	// fmt.Sprintf is good way convert to string
	precision := fmt.Sprintf("%.2f", b)

	// print value and type
	fmt.Printf("%s - %T\n", precision, precision)
}