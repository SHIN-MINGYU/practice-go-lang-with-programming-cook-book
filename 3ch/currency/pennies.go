package currency

import "strconv"

// ConvertPenniesToDollarString function return dollars string expression, after get int64 type pennies amount
func ConvertPenniesToDollarString(amount int64) string {
	// input pennies convert to demical
	result := strconv.FormatInt(amount, 10)

	// if negative number , setting after
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// if pennies amount is less than 100
	// plus "0" at left
    if len(result) < 3 {
		result = "0" + result
	}
	length := len(result)
	//  add to demical
	result = result[:length-2] + "." + result[length-2:]

	// if this is negative, add "-" at left
	if negative {
        result = "-" + result
    }
	return result
}
