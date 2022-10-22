package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// bring string-input from client
		fmt.Printf("Enter some text : ")
		// message input
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading input : %s\n",err.Error())
			continue
		}

		// connect to addr 
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encounted and error connection : %s\n",err.Error())
		}
		// send to server from client
		fmt.Fprintf(conn,data)
		
		// read response again
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reawding response :%s", err.Error())
		}
		fmt.Printf("Received back : %s", status)
		// quit accomplished connection
		conn.Close()
	}
}