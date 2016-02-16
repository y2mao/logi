package logi

import (
	"fmt"
	"os"
)

func FlushAll() {
	for _, lg := range loggerMap {
		lg.persist()
	}
}

func Info(name string, v ...interface{}) {
	writeLog(name, "INFO  | ", fmt.Sprint(v...))
}

func Infof(name string, format string, v ...interface{}) {
	writeLog(name, "INFO  | ", fmt.Sprintf(format, v...))
}

func Warn(name string, v ...interface{}) {
	writeLog(name, "WARN  | ", fmt.Sprint(v...))
}

func Warnf(name string, format string, v ...interface{}) {
	writeLog(name, "WARN  | ", fmt.Sprintf(format, v...))
}

func Error(name string, v ...interface{}) {
	writeLog(name, "ERROR | ", fmt.Sprint(v...))
}

func Errorf(name string, format string, v ...interface{}) {
	writeLog(name, "ERROR | ", fmt.Sprintf(format, v...))
}

func Panic(name string, v ...interface{}) {
	s := fmt.Sprint(v...)
	writeLog(name, "ERROR | PANIC | ", s)
	panic(s)
}

func Panicf(name string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	writeLog(name, "ERROR | PANIC | ", s)
	panic(s)
}

func Fatal(name string, v ...interface{}) {
	writeLog(name, "ERROR | FATAL | ", fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(name string, format string, v ...interface{}) {
	writeLog(name, "ERROR | FATAL | ", fmt.Sprintf(format, v...))
	os.Exit(1)
}
