package logi

import "flag"

var (
	logDir     = flag.String("logi-dir", "", "the directory where log written to")
	quietMode  = flag.Bool("logi-quiet", false, "diable console output. Please turn it on in production env")
	rolling    = flag.String("logi-rolling", "20060102", "rolling format. Please refer to standard RFC time format")
	bufferSize = flag.Int("logi-bufsize", 1024*1024*4, "buffer size for caching log")
	interval   = flag.Int("logi-interval", 15, "flush interval(sec)")
)

func CfgQuietMode() bool {
	return *quietMode
}

func CfgRolling() string {
	return *rolling
}

func CfgLogDir() string {
	if s := *logDir; len(s) > 0 {
		return s
	}

	return "./logs"
}

func CfgBufferSize() int {
	if d := *bufferSize; d > 1024 {
		return d
	}

	return 1024 * 1024 * 4
}

func CfgFlushInterval() int {
	if d := *interval; d > 0 {
		return d
	}

	return 15
}
