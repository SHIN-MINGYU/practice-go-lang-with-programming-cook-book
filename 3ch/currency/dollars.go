package currency

import (
	"errors"
	"strconv"
	"strings"
)

// ConvertStringDollarsToPennies function get dollars amount to string , and convert  to int64 type

func ConvertStringDollarsToPennies(amount string) (int64, error) {
    // Check amount parameter can convert float
	_, err := strconv.ParseFloat(amount, 64)
	if err!= nil {
        return 0, err
    }

	// split string based on '.' character
	group := strings.Split(amount, ".")

	// if '.' character is not exist, value is saved in result
	result := group[0]

	// basic string
	r := ""

	// process data after '.' character
	if len(group) == 2{
		if len(group[1]) != 2 {
			return 0, errors.New("invalid data format")
		}
		r = group[1]
	}

	// fill 0
	// if '.' character is not exist, 0 is filled twice
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// convert to int
	return strconv.ParseInt(result, 10, 64)
}
