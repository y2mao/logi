package logi

import (
	"fmt"
	"os"
)

func logf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stdout, "[logi] "+f+"\n", v...)
}

func panicf(f string, v ...interface{}) {
	logf(f, v...)
	panic(fmt.Sprintf(f, v...))
}
