package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer function show the way how to initailize bytes buffer
// implements io.Reader interface
func Buffer(rawString string) *bytes.Buffer{
	// start at encoded bytes by raw bytes
	rawBytes := []byte(rawString)

	// there are many way to create buffer at raw bytes or original string
	b := new(bytes.Buffer)
	b.Write(rawBytes)

	// another way
	b = bytes.NewBuffer(rawBytes)

	// the way for full avoiding line 9's bytes array
	b = bytes.NewBufferString(rawString)
	return b
}

// ToString function is return string
// after get io.Reader and use all
func toString(r io.Reader) (string , error) {
	b, err := ioutil.ReadAll(r)
	if err != nil{
		return "", err
	}
	return string(b), err
}