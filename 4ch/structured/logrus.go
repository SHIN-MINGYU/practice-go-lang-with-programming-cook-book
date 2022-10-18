package structured

import "github.com/sirupsen/logrus"

// Hook struct is imply interface of hook in logrus
type Hook struct {
	id string
}

// Fire function is called by recording log

func (h *Hook) Fire(entry *logrus.Entry) error {
    entry.Data["id"] = h.id
    return nil
}

// Levels function is level to excute sent hook

func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus function show basic function of logrus

func Logrus(){
	// record log to json format
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When string
	}{"Somthing happened","Just Now"}

	x := logrus.WithFields(fields)
	x.Warn("Warning!")
	x.Error("Error!")
}