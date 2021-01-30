package glog

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Pname = "glog"

func init() {
	path, _ := os.Getwd()
	fmt.Println("当前路径:" + path)
}

func Basef(level string, format string) string {
	ft := "[%v]%v[%v] "
	if len(level) <= 4 {
		ft = "[%v] %v[%v] "
	}
	_, f, l, _ := runtime.Caller(2)
	f = f[strings.LastIndex(f, Pname)+len(Pname)+1:]
	prefix := fmt.Sprintf(ft, level, time.Now().Format("2006-01-02 15:04:05"), f+":"+strconv.Itoa(l))
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
	ft := ""
	if len(level) <= 4 {
		ft = "[%v] %v[%v] " + va
	} else {
		ft = "[%v]%v[%v] " + va
	}
	_, f, l, _ := runtime.Caller(2)
	f = f[strings.LastIndex(f, Pname)+len(Pname)+1:]
	prefix := fmt.Sprintf(ft, level, time.Now().Format("2006-01-02 15:04:05"), f+":"+strconv.Itoa(l))
	return prefix
}

func Debugf(format string, a ...interface{}) {
	format = Basef("DEBUG", format)
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
	format = Basef("ERROR", format)
	color.Red(format, a...)
}

func Debug(a ...interface{}) {
	format := Base("DEBUG", a...)
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
	format := Base("ERROR", a...)
	color.Red(format)
}
