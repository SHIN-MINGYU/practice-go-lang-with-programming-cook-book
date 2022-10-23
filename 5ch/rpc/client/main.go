package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/go-practice-width-cook-book/5ch/rpc/tweak"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("error handling : ", err)
	}
	args := tweak.Args {
		String : "this is string should be uppercase and reserved",
		ToUpper: true,
		Reserve: true,
	}
	var result string

	err = client.Call("StringTweaker.Tweak", args, &result)
	if err != nil {
		log.Fatal("client call with erorr : ", err)
	}
	fmt.Printf("the reuslt is : %s", result)
}