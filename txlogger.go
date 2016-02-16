package txlog

import (
	"bytes"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type txLogger struct {
	name string

	// buffer vars
	buf     bytes.Buffer
	bufSize int

	// persist vars
	persisting     bool
	persistLock    sync.Mutex
	persistCounter uint64
}

func newTXLogger(name string) *txLogger {
	tl := new(txLogger)
	tl.name = name
	tl.bufSize = CfgBufferSize()
	tl.persisting = false

	return tl
}

func (tx *txLogger) buffer(b []byte) {
	tx.persistLock.Lock()
	tx.buf.Write(b)
	tx.persistLock.Unlock()

	if tx.buf.Len() >= tx.bufSize {
		tx.persist()
	}
}

func (tx *txLogger) persist() {
	// ignore if persistance is performing or buffer is empty
	if tx.persisting || tx.buf.Len() == 0 {
		return
	}

	// lock & unlock
	tx.persistLock.Lock()
	tx.persisting = true
	defer func() {
		tx.persistLock.Unlock()
		tx.persisting = false
	}()

	// perf monitoring use
	if tx.persistCounter >= 18446744073709551615 {
		tx.persistCounter = 0
	} else {
		tx.persistCounter++
	}

	// build file full path
	fp := ""
	dir := CfgLogDir()
	rollingText := CfgRolling()
	if len(rollingText) > 0 {
		fp = filepath.Join(dir, time.Now().Format(rollingText)+"."+tx.name+".log")
	} else {
		fp = filepath.Join(dir, tx.name+".log")
	}

	// write data from buffer to file
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panicf(`persist log to "%s": %v`, fp, err)
	}
	defer f.Close()
	tx.buf.WriteTo(f)

	// reset buffer
	tx.buf.Reset()
}
