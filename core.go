package logi

import (
	"os"
	"strings"
	"time"
)

var (
	loggerMap = make(map[string]*txLogger)

	flushTicker *time.Ticker
)

func flushTimely() {
	for t := range flushTicker.C {
		t.String()
		FlushAll()
	}
}

func writeLog(name string, prefix string, s string) {
	// get current time early.
	now := time.Now()

	// retrieve/new logger
	lg, ok := loggerMap[name]
	if !ok {
		lg = newTXLogger(name)
		loggerMap[name] = lg
	}

	// build log content
	var buf []byte
	buf = buf[:0]
	formatHeader(&buf, now)
	buf = append(buf, prefix...)

	//  remove newline char and write content
	buf = append(buf, strings.Replace(s, "\n", "\\n", -1)...)

	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	// write console
	if !CfgQuietMode() {
		os.Stdout.Write(buf)
	}

	// write to in-memory buffer
	lg.buffer(buf)
}

// Folk from "log" package
func formatHeader(buf *[]byte, t time.Time) {
	// ignore
	// *buf = append(*buf, l.prefix...)
	year, month, day := t.Date()
	itoa(buf, year, 4)
	*buf = append(*buf, '/')
	itoa(buf, int(month), 2)
	*buf = append(*buf, '/')
	itoa(buf, day, 2)
	*buf = append(*buf, ' ')

	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)
	// *buf = append(*buf, '.')
	// itoa(buf, t.Nanosecond()/1e3, 6)
	*buf = append(*buf, ' ')
}

// Folk from "log" package
// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
