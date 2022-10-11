package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// Base64Example function show demo to use base64 package
func Base64Example() error {
	// base64 package is useful when, cant support binary format working bytes/string
	// use helper function and URL encoding
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding", value)

	// decode first value
	decoded, err := base64.URLEncoding.DecodeString(value)
    if err!= nil {
        return err
    }
	fmt.Println("With DecodeToString and URLEncoding", string(decoded))
	return nil
}

// Bse64ExampleEncoder function show similar example to use encoder / decoder

func Base64ExampleEncoder() error {
    //use encoder/decoder
	buffer := bytes.Buffer{}
	
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)
	
	if _,err := encoder.Write([]byte("encoding some other data!")); err != nil {
		return err
	}

	// Check for closing encoder
	if err := encoder.Close(); err!= nil {
        return err
    }

	fmt.Println("Using encoder and StdEncoding : ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}

	fmt.Println("Using decoder and StdEncoding : ", string(results))

	return nil
}
