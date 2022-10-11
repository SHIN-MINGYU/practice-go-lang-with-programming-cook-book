package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type nullInt64 sql.NullInt64

// ExmapleNullInt structure is same
// the difference is that use sql.NullInt64
type ExampleNullInt struct {
	Age *nullInt64 `json:"age,omitempty"`
	Name string `json:"name"`
}

func (v *nullInt64) MarsharlJson() ([]byte, error){
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

// NullEncoding func show the other way How to handle nil/ommited value
func NullEncoding() error {
	e := ExampleNullInt{}

	// no mean invaild value
	if err := json.Unmarshal([]byte(jsonBlob),&e); err != nil {
		return err 
	}
	fmt.Printf("nullInt64 Unmarshal, no age : %+v\n",e)

	value ,err := json.Marshal(&e)
	if err!= nil {
        return err
    }
	fmt.Println("nullInt64 Marshal, with no age :", string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err!= nil {
		return err
	}

	fmt.Printf("nullInt64 Unmarshal, with age : %+v\n",e)

	value,err = json.Marshal(&e)
    if err!= nil {
		return err 
	}
	fmt.Println("nullInt64 Marshal, with age = 0", string(value))
	return nil
}

