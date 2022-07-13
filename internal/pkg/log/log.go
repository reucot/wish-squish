package log

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/reucot/wish-squish/config"
	"github.com/sirupsen/logrus"
)

var LogOutput *os.File

type Entry struct {
	*logrus.Entry
}

func ReloadLog() {
	CloseLog()
	OpenLog()
}

func OpenLog() error {
	logfile, err := filepath.Abs(filepath.Join("../..", config.Get().Log.Output))

	if err != nil {
		return err
	}

	LogOutput, err = os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	logrus.SetOutput(LogOutput)

	return nil
}

func CloseLog() error {
	if LogOutput != nil {
		err := LogOutput.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func Error(message string, v ...interface{}) {
	logrus.Error(fmt.Sprintf(message, v...))
}
