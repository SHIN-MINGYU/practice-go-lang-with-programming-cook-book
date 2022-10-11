package collections

// WorkWith struct is using for implement collection

type WorkWith struct {
    Data string
	Version int
}

// Filter function is functional filter, this function return restaurant list to input WorkWith function return each "true" property of bool, WorkWith list
func Filter(ws []WorkWith, f func(w WorkWith) bool)[]WorkWith {
	// According to result , reduce the case of result's length 0
	result := make([]WorkWith, 0)
	for _, w := range ws {
        if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map function is functional map, After input WorkWith list and WorkWith, 
// get parameter what function return modifyed WorkWith
// fianl result is modifyed WorkWith list

func Map(ws []WorkWith, f func(WorkWith) WorkWith) []WorkWith {
	// result always have same length
    result := make([]WorkWith, len(ws))
    for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
    }
	return result
}
