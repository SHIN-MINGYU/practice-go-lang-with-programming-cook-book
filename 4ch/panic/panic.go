package panic

import (
	"fmt"
	"strconv"
)

// Panic function is occured error by divide from 0
func Panic(){
	zero, err := strconv.ParseInt("0",10,64)
	if err != nil {
		panic(err)
	}
	a := 1 / zero
	fmt.Println("we will not see this message", a)
} 


// Catcher fucntion call Panic function

func Catcher(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("panic occured : ",r)
		}
	}()
	Panic()
}