package main

// go lang's strong interface -> io.Reader, io.Writer

// implemention of the interface
type Reader interface {
	Read(p []byte)(n int, err error)
}
type Writer interface {
	Write(p []byte)(n int,err error)
}

// go lang offer way to combine interfaces

type Seeker interface {
	Seek(offset int64, whence int)(int64 ,error)
}

type ReadSeeker interface {
	Reader
	Seeker
}

func main() {
}
