package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	// lock for map's modification
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections){
	count := 0

	for {
		count++
		conns.mu.Lock()

		// loop about confirmed address
		for _, retAddr := range conns.addrs {
			// send message to everywhere
			msg := fmt.Sprintf("Send %d",count)
			if _ , err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered : %s", err.Error())
				continue
			}
		}
		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}