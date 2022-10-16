package log

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError function return original error
func OriginalError() error {
	return errors.New("An error occured")
}

// PassThroughError function call OriginalError function and wrapping that, send error

func PassThroughError() error {
	err := OriginalError()

	return errors.Wrap(err, "in passthroughError")
}


// FinalDestination function handle error but, dont represent error

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// unexpected somthing is occured
		// so start logging
		log.Printf("an error occured : %s\n", err.Error())
		return
	}
}

