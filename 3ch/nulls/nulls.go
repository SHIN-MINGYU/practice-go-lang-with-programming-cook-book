package nulls

import (
	"encoding/json"
	"fmt"
)

// the json of name is exist but age is not exist
const (
	jsonBlob = `{"name":"Mingyu"}`
	fullJsonBlob = `{"name":"Mingyu","age":0}`
)

// Example structure have age and name fields as keys
type Example struct {
    Name string `json:"name"`
    Age  int  `json:"age, omitempty"`
}

// BaseEncoding function show how to incode and decode about normal type
func BaseEncoding() error{
	e := Example{}

	// caution!
	// no age is 0 age
	if err := json.Unmarshal([]byte(jsonBlob),&e); err != nil{
		return err
	}

	fmt.Printf("Regular Unmarshal, no age : %+v\n",e)
	value, err := json.Marshal(&e)
	if err!= nil{
        return err
    }
	fmt.Println("Regular Marshal, with age = 0",string(value))
	return nil
}