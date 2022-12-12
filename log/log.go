package log

import (
	"fmt"
	"io"
	"os"
)

var (
	writer io.Writer
	format = JSONFormat
)

func init() {
	writer = io.MultiWriter(os.Stdout)
}

func SetType(logFormat int) {
	format = logFormat
}

func SetWriter(writers ...io.Writer) {
	writer = io.MultiWriter(writers...)
}

func Debug(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}
	data := newData(LevelDebug, message)
	fmt.Fprintln(writer, fmt.Sprintf("%s", data))
}

func Info(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelInfo, message)
	fmt.Fprintln(writer, fmt.Sprintf("%s", data))
}

func Warning(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelWarning, message)
	fmt.Fprintln(writer, fmt.Sprintf("%s", data))
}

func Error(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelError, message)
	fmt.Fprintln(writer, fmt.Sprintf("%s", data))
}
