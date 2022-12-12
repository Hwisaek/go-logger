package log

import (
	"encoding/json"
	"runtime"
	"time"
)

type data struct {
	time    time.Time
	level   int
	file    string
	line    int
	message string
}

func newData(level int, message string) data {
	_, file, line, _ := runtime.Caller(1)
	d := data{
		time:    time.Now(),
		level:   level,
		file:    file,
		line:    line,
		message: message,
	}
	return d
}

func (r data) String() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}
