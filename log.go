package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type LogLevelValue uint8

const (
	TRACE LogLevelValue = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames []string = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

}
func (p *LogLevelValue) FromString(s string) {
	for i, v := range levelNames {
		if strings.EqualFold(v, s) {
			*p = LogLevelValue(i)
			return
		}
	}
	*p = INFO

}

func (p *LogLevelValue) String() string {
	return levelNames[*p]
}

// Stores the log level to use. The default is ERROR
var LogLevel LogLevelValue = ERROR

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}

func Tracef(str string, v ...any) {
	logF(TRACE, str, v...)
}

func Traceln(v ...any) {
	logN(TRACE, v...)
}
func Debugf(str string, v ...any) {
	logF(DEBUG, str, v...)
}

func Debugln(v ...any) {
	logN(DEBUG, v...)
}
func Infof(str string, v ...any) {
	logF(INFO, str, v...)
}

func Infoln(v ...any) {
	logN(INFO, v...)
}

func Warnf(str string, v ...any) {
	logF(WARN, str, v...)
}

func Warnln(v ...any) {
	logN(WARN, v...)
}
func Fatal(v ...any) {
	log.Fatal(v...)
}

func Errorf(str string, v ...any) {
	logF(ERROR, str, v...)
}

func Errorln(v ...any) {
	logN(ERROR, v...)
}

func logF(level LogLevelValue, str string, v ...any) {
	pc, _, _, ok := runtime.Caller(2)
	go _logF(pc, ok, level, str, LogLevel, v...)
}

func _funcName(pc uintptr) string {
	name := runtime.FuncForPC(pc).Name()
	return name[strings.LastIndex(name, "/")+1:]
}

func _logF(pc uintptr, ok bool, level LogLevelValue, str string, currentLogLevel LogLevelValue, v ...any) {
	var color int
	var name string
	if ok {
		name = _funcName(pc)
	}

	switch level {
	case TRACE:
		color = Cyan
	case ERROR, FATAL:
		color = Red
	case DEBUG:
		color = Magenta
	case WARN:
		color = Yellow
	case INFO:
		color = Green

	}

	if level >= currentLogLevel {
		log.Printf(Colourize(fmt.Sprintf("%-5s [%s] %s", levelNames[level], name, str), color), v...)
	}
}

func logN(level LogLevelValue, v ...any) {
	pc, _, _, ok := runtime.Caller(2)
	go _logN(pc, ok, level, LogLevel, v...)
}

func _logN(pc uintptr, ok bool, level LogLevelValue, currentLogLevel LogLevelValue, v ...any) {
	var color int
	var name string
	if ok {
		name = _funcName(pc)
	}

	switch level {
	case TRACE:
		color = Cyan
	case ERROR, FATAL:
		color = Red
	case DEBUG:
		color = Magenta
	case WARN:
		color = Yellow
	case INFO:
		color = Green

	}
	if level >= currentLogLevel {
		kk := []interface{}{fmt.Sprintf("%-5s ", levelNames[level]), fmt.Sprintf("[%s] ", name)}

		kk = append(kk, v...)
		println(color, kk...)
	}
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func println(color int, v ...interface{}) {
	log.Println(Colourize(fmt.Sprint(v...), color))
}

func Println(v ...any) {
	logN(INFO, v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...any) {
	logF(INFO, format, v...)
}
