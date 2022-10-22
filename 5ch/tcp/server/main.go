package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn){
	// set Reader in conn
	reader := bufio.NewReader(conn)

	// bring brought-data's first line 
	data ,err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data : %s\n", err.Error())
		return
	}

	// send data after printing
	fmt.Printf("Received : %s", data)
	conn.Write([]byte(strings.ToUpper(data)))
	
	// close accomplished connection
	conn.Close()
}

func main(){
	In, err := net.Listen("tcp",addr)
	if err != nil {
		panic(err)
	}

	defer In.Close()
	fmt.Printf("listening on %s\n",addr)

	for {
		// accept data from addr connected client by tcp protocol
		conn, err := In.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting connection : %s\n", err.Error())
			// if error exist, retry 
			continue
		}

		// this will be good use-case for worker pool potentially
		go echoBackCapitalized(conn)
	}
}

