package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader convert http connection to websocket
// here use standard buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
}


func wsHandler(w http.ResponseWriter, r *http.Request){
	// upgrade connection
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println("failed to upgrade websocket connection : ", err)
		return
	}
	for {
		// send same message after reading message
		messageType, p,err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message : ",err)
			return 
		}
		log.Printf("received from client :%#v", string(p))
		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println("failed to write message : ",err)
			return
		}
	}
}
