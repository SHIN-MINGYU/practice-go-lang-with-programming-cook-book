package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-practice-width-cook-book/2ch/envvar"
)

// Config struct for saving configuration read by environment variables and json files
type Config struct {
	Version string `json : "version" required : "true"`
	IsSafe bool `json : "is_safe" default : "true"` 
	Secret string `json : "secret"`
}

func main() {
	var err error
	
	// To save save example json file
	// create temporary file
	tf, err := ioutil.TempFile("","tmp")
	if err != nil {
		panic(err)
	}

	defer tf.Close()
	defer os.Remove(tf.Name())
	
	// To save secret value
	// create json file

	secrets := `{
		"secret" : "so so secret"
		}`
	if _,err := tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}
	
	// If you need, can set environment variables easily
	if err := os.Setenv("EXAMPLE_VERSION","1.0.0"); err != nil {
		panic(err)
	}
	if err := os.Setenv("EXAMPLE_IS_SAFE","false"); err!= nil {
		panic(err)
	}
	c := Config{}
	if err := envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}
	fmt.Println("secrets file contains =", secrets)

	fmt.Println("GO_PATH = ",os.Getenv("GOPATH"))
	// also easily read environment variables
	fmt.Println("EXAMPLE_VERSION : ",os.Getenv(("EXAMPLE_VERSION")))
	fmt.Println("EXAMPLE_IS_SAFE : ",os.Getenv(("EXAMPLE_IS_SAFE")))

	// final config is mixed json file and environment variables
	fmt.Printf("Final config : %#v\n",c)
}