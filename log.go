package glog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func Basef(level string, format string) string {
	ft := "[%v]%v[%v] "
	_, f, l, _ := runtime.Caller(2)
	path, _ := os.Getwd()
	f = f[len(path)+1:]
	prefix := fmt.Sprintf(ft, level, time.Now().Format("2006-01-02/15:04:05"), f+":"+strconv.Itoa(l))
	return prefix + format
}

func Base(level string, args ...interface{}) string {
	format := ""
	lt := len(args)
	if lt > 0 {
		format = "%v"
	}
	for i := 1; i < lt; i++ {
		format += " %v"
	}
	va := fmt.Sprintf(format, args...)
	ft := "[%v]%v[%v] " + va
	path, _ := os.Getwd()
	_, f, l, _ := runtime.Caller(2)
	f = f[len(path)+1:]
	prefix := fmt.Sprintf(ft, level, time.Now().Format("2006-01-02/15:04:05"), f+":"+strconv.Itoa(l))
	return prefix
}

func Debugf(format string, a ...interface{}) {
	format = Basef("DEBU", format)
	color.Blue(format, a...)
}

func Infof(format string, a ...interface{}) {
	format = Basef("INFO", format)
	color.Cyan(format, a...)
}

func Warnf(format string, a ...interface{}) {
	format = Basef("WARN", format)
	color.Yellow(format, a...)
}

func Errorf(format string, a ...interface{}) {
	format = Basef("ERRO", format)
	color.Red(format, a...)
}

func Debug(a ...interface{}) {
	format := Base("DEBU", a...)
	color.Blue(format)
}

func Info(a ...interface{}) {
	format := Base("INFO", a...)
	color.Cyan(format)
}

func Warn(a ...interface{}) {
	format := Base("WARN", a...)
	color.Yellow(format)
}

func Error(a ...interface{}) {
	format := Base("ERRO", a...)
	color.Red(format)
}
