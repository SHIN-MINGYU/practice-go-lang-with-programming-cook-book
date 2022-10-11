package tags

import (
	"reflect"
)

// SerializeStructStrings function convert user sutom serialization format from struct
// it use the way serialize struct tag about type of string

func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// convert interface to spefic type
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)


	// if pointer about struct is sent, do something
	if r.Kind() == reflect.Ptr {
        r = r.Elem()
        value = value.Elem()
    }

	// do loop about all fields
	for i := 0; i < r.NumField(); i++ {
        field := r.Field(i)
		// if struct tag is discovered,
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignore "-", Otherwise all value be serialize key
			if serialize == "-" {
                continue
            }
			key = serialize
		}

		switch value.Field(i).Kind(){
			// this example support just string
			case reflect.String:
                result += key + ":" + value.Field(i).String() + ";"
			default :
				continue
		}
	}
	return result, nil
}