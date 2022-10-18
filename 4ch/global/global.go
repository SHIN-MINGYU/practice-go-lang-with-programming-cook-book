package global

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// create global-package-level-variables tp LowerCase
var (
	log *logrus.Logger
	initLog sync.Once
)

// Init function set logger at first and return error when excute more than twice
func Init() error {
	err := errors.New("already initialized")
	 initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel	
	})
	return err
}

// SetLog function set log
func SetLog(l *logrus.Logger) {
	log = l
}

// WithField function send connected-log at WithField to global log
func WithField(key string, value interface{}) *logrus.Entry {
    return log.WithField(key, value)
}

// Debug function send connected-log at Debug to global log
func Debug(args ...interface{}){
	log.Debug(args...)
}
