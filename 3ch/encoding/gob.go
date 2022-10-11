package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// pos struct is saved x, y position

type pos struct {
    X int
    Y int
	Object string
}

// GobExample function show How to use gob package

func GobExample() error {
	buffer := bytes.Buffer{}

    // create object
    p := pos{
        X: 10,
        Y: 20,
        Object: "Hello world",
    }

    // if p is interface , call gob.Register first

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	fmt.Println("Gob Encoded valued len: ", len(buffer.Bytes()))

	p2 := pos{}

	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err!= nil {
		return err 
	}

	fmt.Println("Gob Decode value : ", p2)
	return nil
}
