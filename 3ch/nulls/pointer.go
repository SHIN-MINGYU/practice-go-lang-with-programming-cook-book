package nulls

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer structure is same(nulls.go's Example structure),
// the diffrence is using *Int
type ExamplePointer struct {
	Age *int `json:"age, omitempty"`
	Name string `json:"name"`
}

// Pointer Encoding function show How to process nil/ommited value
func PointerEncoding() error {
	// no age is nil age
	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonBlob),&e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, no age : %+v\n",e)

	value, err := json.Marshal(&e)
	if err!= nil {
        return err
    }
	fmt.Println("Pointer Marshal, with no age : ",string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, with age = 0 : %+v\n",e)

	value, err = json.Marshal(&e)
	if err!= nil {
        return err
    }
	fmt.Println("Pointer Marshal, with age = 0 : ",string(value))

	return nil
	
}