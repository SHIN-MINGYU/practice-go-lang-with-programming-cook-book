package collections

import "strings"

// LowerCaseData function conduct ToLower in WorkWith's Data property

func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion function increase WorkWith's Version property

func IncrementVersion(w WorkWith) WorkWith {
    w.Version++
    return w
}

// OldVersion function return closure the version is greater than specified version

func OldVersion(v int) func(w WorkWith) bool {
    return func(w WorkWith) bool {
		return w.Version >= v
	}
}
