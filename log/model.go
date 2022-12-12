package log

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type data struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Source  string    `json:"source"`
	Message string    `json:"message"`
}

func newData(level, message string) data {
	_, file, line, _ := runtime.Caller(2)
	d := data{
		Time:    time.Now(),
		Level:   level,
		Source:  fmt.Sprintf("%s:%d", file, line),
		Message: message,
	}
	return d
}

func (r data) String() string {
	msg := ""
	switch format {
	case NormalFormat:
		marshal, _ := json.Marshal(r)
		msg = string(marshal)
	default:
		msg = fmt.Sprintf("%v %v %v %v", r.Time.Format("2006-01-02 15-04-05.000"), r.Level, r.Source, r.Message)
	}
	return msg
}
