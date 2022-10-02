package interfaces

import (
	"io"
	"os"
)

// PipeExample function offer more example for using interface
func PipeExample() error{

	// pipe reader and pipe wirter are implement io.Reader and io.Writer
	// this is inmemory pipe
	// it is purpose to read stream at same time in same stream while writing another source 
	r, w := io.Pipe()

	// this code should excute in separate go-routin
	// Because It works in a block way.
	// this is the reason reader wait the moment of shut down for cleaning up 
	go func(){
		// simple content 
		// also can encoding to  json, base64 ...etc
		w.Write([]byte("text\n"));
		w.Close()
	}()

	if _,err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil;
}