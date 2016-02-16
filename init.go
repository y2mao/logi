package logi

import (
	"flag"
	"os"
	"time"
)

func init() {
	// parse variables from args
	flag.Parse()

	// create folder if not exist
	if _, err := os.Stat(CfgLogDir()); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(CfgLogDir(), 0777)
	}

	// start ticker for interval flush
	flushTicker = time.NewTicker(time.Duration(CfgFlushInterval()) * time.Second)
	go flushTimely()
}
