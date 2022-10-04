package filedirs

import (
	"errors"
	"io"
	"os"
)

// Operate function is to control file or directory
func Operate() error {
	// this 0755 is similar Command Line's Chown
	// this command create /tmp/example dir
	// also, instead of relative path, can use absolute path
	if err := os.Mkdir("example_dir",os.FileMode(0755));
	err != nil {
		return err 
	}
	// cd /tmp dir
	if err := os.Chdir("example_dir",); err != nil {
		return err 
	}

	// f is normal file object
	// implements some interface
	// When open the file , if set appropriate bit,
	// can use immediately as reader or writer
	f, err := os.Create("test.txt");
	if err != nil {
		return err
	}

	// After write known-length in file,
	// verify the result to exactly complete
	value := []byte("hello");
	count, err := f.Write(value);
	if err != nil {
		return err
	}
	if count != len(value){
		return errors.New("incorrect length returned from write")
	}
	if err := f.Close(); err != nil {
		return err
	}

	// read file
	f, err = os.Open("test.txt");
	
	if err != nil {
		return err 
	}

	io.Copy(os.Stdout,f)

	if err := f.Close(); err != nil {
		return err 
	}

	// cd /tmp directory
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// if the case to indicate wrong directory or use user input
	// cleanup, os.RemoveAll will be dangerous 
	// specially, root excute is dangerous too

	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}
	return nil
}