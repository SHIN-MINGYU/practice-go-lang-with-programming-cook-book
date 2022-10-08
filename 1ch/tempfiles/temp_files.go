package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp function for temp files and directory processing
func WorkWithTemp() error {
	// For example, if you need space for saving temporary files what have same name
	// temp directory is good way to access

	// first argument is blank this means that create directory at returned place by os.TempDir()
	t, err := ioutil.TempDir("","tmp")
	if err!= nil {
        return err
    }

	// when conduct this work later
	// temp files all content is deleted when function is ended
	defer os.RemoveAll(t)

	// directory must be existed for creating temporary files
	// tf is *os.File Object
	tf, err := ioutil.TempFile(t, "tmp")
	if err!= nil {
        return err
    }

	fmt.Println(tf.Name())

	// Usually delete temporary files at here
	// However, it is clear by defet statements why save files in temporary directory
	return nil
}
