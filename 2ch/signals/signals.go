package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig function set listner about SIGINT interupt
func CatchSig(ch chan os.Signal, done chan bool) {
	// block for wait signal
	sig := <- ch

	// if signal is already, print
	fmt.Println("\n Sig returned: ", sig)

	// can set handler about all signal
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT signal")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM signal")
	default:
		fmt.Println("unexpected signal returned")
	}
	done <- true
}

func main() {
    // make channel
	ch := make(chan os.Signal)
	done := make(chan bool)
    
	// regist to signals library
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// if got the signal, write to done
	go CatchSig(ch, done)

	fmt.Println("Press Ctrl+C to terminate...")
	<-done
	fmt.Println("Done")
}
