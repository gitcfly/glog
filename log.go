package glog

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	DEBUG string = "DEBUG"
	INFO  string = "INFO"
	WARN  string = "WARN"
	ERROR string = "ERROR"
	FATAL string = "FATAL"
)

var LNumb = map[string]int{
	DEBUG: 1,
	INFO:  2,
	WARN:  3,
	ERROR: 4,
	FATAL: 5,
}

var Conf = &Cfg{Level: DEBUG, LevCut: true}

type Cfg struct {
	PName  string
	Level  string
	LevCut bool
}

func CfgLog(cfg *Cfg) {
	Conf = cfg
}

func BaseFormat(level string, format string) (string, bool) {
	if LNumb[level] < LNumb[Conf.Level] {
		return "", false
	}
	if Conf.LevCut {
		level = level[:4]
	}
	_, fe, l, _ := runtime.Caller(2)
	idx := strings.Index(fe, Conf.PName)
	fie := fe[idx+len(Conf.PName)+1:] + ":" + strconv.Itoa(l)
	tmef := time.Now().Format("2006-01-02/15:04:05")
	prefix := fmt.Sprintf("[%v]%v[%v] ", level, tmef, fie)
	return prefix + format, true
}

func Base(level string, args ...interface{}) (string, bool) {
	if LNumb[level] < LNumb[Conf.Level] {
		return "", false
	}
	if Conf.LevCut {
		level = level[:4]
	}
	var format = strings.Repeat(" %v", len(args))
	format = strings.TrimLeft(format, " ")
	va := fmt.Sprintf(format, args...)
	ft := "[%v]%v[%v] " + va
	_, f, l, _ := runtime.Caller(2)
	idx := strings.Index(f, Conf.PName)
	fe := f[idx+len(Conf.PName)+1:] + ":" + strconv.Itoa(l)
	tmef := time.Now().Format("2006-01-02/15:04:05")
	prefix := fmt.Sprintf(ft, level, tmef, fe)
	return prefix, true
}

func Debugf(format string, a ...interface{}) {
	if lgf, ok := BaseFormat(DEBUG, format); ok {
		color.Blue(lgf, a...)
	}
}

func Infof(format string, a ...interface{}) {
	if lgf, ok := BaseFormat(INFO, format); ok {
		color.Cyan(lgf, a...)
	}
}

func Warnf(format string, a ...interface{}) {
	if lgf, ok := BaseFormat(WARN, format); ok {
		color.Yellow(lgf, a...)
	}
}

func Errorf(format string, a ...interface{}) {
	if lgf, ok := BaseFormat(ERROR, format); ok {
		color.Red(lgf, a...)
	}
}

func Fatalf(format string, a ...interface{}) {
	if lgf, ok := BaseFormat(DEBUG, format); ok {
		color.Magenta(lgf, a...)
	}
}

func Debug(a ...interface{}) {
	if lgf, ok := Base(DEBUG, a...); ok {
		color.Blue(lgf)
	}
}

func Info(a ...interface{}) {
	if lgf, ok := Base(INFO, a...); ok {
		color.Cyan(lgf)
	}
}

func Warn(a ...interface{}) {
	if lgf, ok := Base(WARN, a...); ok {
		color.Yellow(lgf)
	}
}

func Error(a ...interface{}) {
	if lgf, ok := Base(ERROR, a...); ok {
		color.Red(lgf)
	}
}

func Fatal(a ...interface{}) {
	if lgf, ok := Base(FATAL, a...); ok {
		color.Magenta(lgf)
	}
}
