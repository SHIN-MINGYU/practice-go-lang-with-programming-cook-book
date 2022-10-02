package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString function show many way what search string
func SearchString(){
	s := "this is a test"

	// return true
	// because s include 'test'
	fmt.Println(strings.Contains(s, "this"))

	// return true
	// because s include alphabet 'a'
	// return same result in the case what dont include b or c
	fmt.Println(strings.ContainsAny(s,"abc"))

	// return true
	// because s is started 'this'
	fmt.Println(strings.HasPrefix(s, "this"))

	// return true 
	// because s is ended 'test'
	fmt.Println(strings.HasSuffix(s, "this"))
}

// ModifyString function modify string in variable way
func ModifyString(){
	s := "simple string"

	// print [simple string]
	fmt.Println(strings.Split(s, " "))

	// print "Simple String"
	fmt.Println(strings.Title(s))

	// print "simple string"
	// remove back and forth white space
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}


// StringReader function show the way what create io.Reader interface to string
func StringReader(){
	s := "simple string"
	r := strings.NewReader(s)

	// print at stdout
	io.Copy(os.Stdout, r)
}