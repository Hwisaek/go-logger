package log

import (
	"encoding/json"
	"runtime"
	"time"
)

type data struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	File    string    `json:"file"`
	Line    int       `json:"line"`
	Message string    `json:"message"`
}

func newData(level, message string) data {
	_, file, line, _ := runtime.Caller(1)
	d := data{
		Time:    time.Now(),
		Level:   level,
		File:    file,
		Line:    line,
		Message: message,
	}
	return d
}

func (r data) String() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}
