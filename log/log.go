package log

import (
	"fmt"
	"log"
)

func Debug(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}
	data := newData(LevelDebug, message)
	log.Println(fmt.Sprintf("%s", data))
}

func Info(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelInfo, message)
	log.Println(fmt.Sprintf("%s", data))
}

func Warning(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelWarning, message)
	log.Println(fmt.Sprintf("%s", data))
}

func Error(messages ...interface{}) {
	message := ""
	for _, m := range messages {
		message += fmt.Sprintf("%v", m)
	}

	data := newData(LevelError, message)
	log.Println(fmt.Sprintf("%s", data))
}
