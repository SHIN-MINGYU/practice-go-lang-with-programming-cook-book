package tweak

import "strings"

// StringTweaker is string type can reverse itself
type StringTweaker struct {}

// Args is option of way for mediate string
type Args struct {
	String string
	ToUpper bool
	Reserve bool
}

// Tweak function comply rpc library complied next content
// - export function's type
// - export function
// - function have two parameter, both is exported type
// - function's second parameter is pointer
// - function have error return type 
func (s StringTweaker) Tweak(args *Args, resp *string) error {

	result := string(args.String)
	if args.ToUpper {
		result = strings.ToUpper(result)
	}

	if args.Reserve {
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i<j; i, j = i+1, j-1 {
			runes[i] , runes[j] = runes[j], runes[i]
		}
		result = string(runes)
 	}
	*resp = result
	return nil
}