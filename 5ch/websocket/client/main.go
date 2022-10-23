package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

// catchSig function make out websocket connection if the program is end by ctrl-c key
func catchSig(ch chan os.Signal, c *websocket.Conn){
	// blocking for wait os.Signal 
	<-ch
	err := c.WriteMessage(websocket.CloseMessage,websocket.FormatCloseMessage(websocket.CloseNormalClosure,""))
	if err != nil {
		log.Println("write close : ",err)
		return
	}
}

func main(){
	// connect os.Signal to chan
	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt)

	// use ws:// struct for connecting websocket
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u,nil)
	if err != nil {
		log.Fatal("dial : ", err)
	}
	defer c.Close()

	// send to catchSig
	go catchSig(interrupt,c)

	process(c)
}