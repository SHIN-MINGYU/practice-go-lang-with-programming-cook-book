package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer function will use buffer created by Buffer function
func WorkWithBuffer() error{
	rawString := "it's easy to encode unicode into a byte array";
	b := Buffer(rawString)

	// convert from Buffer to bytes by using b.Bytes() 
	// or from bytes to string by using b.String()

	fmt.Println(b.String());

	// Because it is io.Reader, 
	// we can use normal io reader's interface function
	
	s, err := toString(b)
	if err != nil {
		return nil;
	}

	fmt.Println(s)

	// also we can create "bytes reader" by bring bytes
	// this readers implement io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, io.RuneScanner interface 
	reader := bytes.NewReader([]byte(rawString))

	// this can connect to scanner what allow tokenzation and reading bufferd data
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// proccess all scanning event to loop
	for scanner.Scan(){
		fmt.Print(scanner.Text())
	}
	return nil
}