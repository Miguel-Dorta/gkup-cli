package log

import (
	"fmt"
	"github.com/Miguel-Dorta/logolang"
	"os"
)

var l *logolang.Logger

func init() {
	l = logolang.NewLogger()
	l.Color = true
	l.Formatter = func(levelName, msg string) string {
		return fmt.Sprintf("[%s] %s", levelName, msg)
	}
	l.Level = logolang.LevelDebug
}

func Critical(message string) {
	l.Critical(message)
	os.Exit(1)
}

func Criticalf(format string, v ...interface{}) {
	l.Criticalf(format, v...)
	os.Exit(1)
}

func Error(message string) {
	l.Error(message)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Info(message string) {
	l.Info(message)
}

func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func Debug(message string) {
	l.Debug(message)
}

func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}
