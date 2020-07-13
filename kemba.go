package kemba

import (
	"fmt"
	"os"
)

type Logger struct {
	tag     string
	filter  interface{}
	allowed string
}

func New(t string) Logger {
	return Logger{tag: t, allowed: os.Getenv("DEBUG")}
}

func (l Logger) Printf(format string, v ...interface{}) {
	if l.allowed == "" {
		return
	}
	fmt.Printf(format, v...)
}
