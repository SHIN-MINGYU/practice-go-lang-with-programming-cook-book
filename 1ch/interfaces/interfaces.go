package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy function copy data from in to out's first directory by using Buffer
// write data too in standard output ( stdout )
func Copy(in io.ReadSeeker, out io.Writer) error {

	// write data parameter 'out' and standard output
	w := io.MultiWriter(out, os.Stdout)


	// if standart copy and parameter 'in' have mass data, this way is dangerous
	if  _,err := io.Copy(w,in); err != nil{
		return err
	}

	in.Seek(0,0)

	//write data to buffer by using 64byte chunk
	buf := make([]byte,64)

	if _,err := io.CopyBuffer(w,in,buf); err != nil {
		return err
	}
	// print at new command line
	fmt.Println()
	return nil
}