package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {
	fmt.Printf("client for server url : %s\n", addr)
	addr, err := net.ResolveUDPAddr("udp",addr)
	if err != nil {
		panic(err)
	}

	conn,err := net.DialUDP("udp",nil,addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := make([]byte,512)
	_, err = conn.Write([]byte("connected"))
	if err != nil {
		panic(err)
	}

	for {
		_, err = conn.Read(msg)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n",msg)
	}
}