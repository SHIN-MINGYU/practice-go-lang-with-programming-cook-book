package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays is custom type to read flag
 type CountTheWays []int
func (c *CountTheWays) String() string {
	result  := ""
	for _,v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Set function is used in flag package
func(c *CountTheWays) Set(v string) error{
	values := strings.Split(v, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
        if err!= nil {
            return err
        }
		*c = append(*c, i)
	}
	return nil
}