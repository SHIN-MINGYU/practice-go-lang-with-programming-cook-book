package structured

import (
	"errors"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// ThrowError function generate error to track
func ThrowError() error{
	err := errors.New("a crazy failure")
	log.WithField("id","123").Trace("ThrowError").Stop(&err)
	return err
}

// CustomHandler structure divide two stream
type CustomHandler struct {
	id string
	handler log.Handler
}

// HandleLog function record log to add hook
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return  h.handler.HandleLog(e)
}

// Apex function provice useful technique 
func Apex(){
	log.SetHandler(&CustomHandler{"123",text.New(os.Stdout)})
	err := ThrowError()

	log.WithError(err).Error("and error occurred")
}