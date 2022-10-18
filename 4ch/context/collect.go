package context

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initailize function call three function for setting
// after then, record log before end
func Initailize(){
	// set basic log
	log.SetHandler(text.New(os.Stdout))
	// initailize context
	ctx := context.Background()
	// create logger and connect to context
	ctx ,e := FromContext(ctx, log.Log)

	// set Fields
	ctx = WithField(ctx,"id","123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gether name")
	gatherLocation(ctx)
	e.Info("after gather location")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx,"name","Go CookBook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx,log.Fields{"city" : "Seattle", "state" : "WA"})
}