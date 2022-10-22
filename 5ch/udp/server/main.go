package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {
	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}
	fmt.Printf("serving on %s\n", addr)

	// create udp addr
	addr, err := net.ResolveUDPAddr("udp",addr)
	if err != nil {
		panic(err)
	}
	
	// wait request to designated address
	conn, err := net.ListenUDP("udp",addr)
	if err != nil {
		panic(err)
	}
	// close connection (delayed call)
	defer conn.Close()

	// send message to all confirmed address by asynchronous
	go broadcast(conn,conns)

	msg := make([]byte, 1024)
	for {
		// receive message for collect port's infomation and address what send message again
		_, retAddr,err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		// save received message in map
		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		fmt.Printf("%s connected\n",retAddr)
 	}
}