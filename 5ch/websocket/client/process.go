package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text : ")
		// because of preventing ctrl-c, so If client want close connection, press ctrl-c and enter
		data, err := reader.ReadString('\n');
		if err != nil {
			log.Println("failed to read stdin ", err)
		}

		// delte space from read data
		data = strings.TrimSpace(data)

		// write byte-typed-message to web socket 
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message ", err)
			return 
		}

		//  can read after write . Because it is echo server
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read : ", err)
			return
		}

		log.Printf("received back from server : %#v\n", string(message))
	}
}