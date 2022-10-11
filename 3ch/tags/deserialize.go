package tags

import (
	"errors"
	"reflect"
	"strings"
)

// DeserializeStructString function convert to struct from serialization string used user custom serialize format
func DeserializeStructString(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	// Always sent the pointer
	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
    }

	// reverse referencing pointer
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// save into map after divide serialized string 
	vals := strings.Split(s,";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}
	//  do loop about all fields 
	for i := 0; i < r.NumField(); i++ {
        field := r.Field(i)
		
		// check how about check going serialized
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
		// ignore "-", Otherwise all value be serialization key
			if serialize == "-" {
                continue
            }
			// check map have serialize variable
			if val, ok := valMap[serialize]; ok {
            	value.Field(i).SetString(val)
			}
			
        }else if val, ok := valMap[field.Name]; ok {
		// check map have field name
			value.Field(i).SetString(val)        
		}
	}
	return nil
}
