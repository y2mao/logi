package txlog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	Info("http", "server started")
	Infof("http", "listening port:%d", 18801)
	Error("api", "undefined API")
	Errorf("api", "%s failed. %v", "/user/profile", errors.New("Invalid ID"))

	FlushAll()
}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Errorf("apicall", "[WorkItemService][UnsafeGetInstanceKey] [WorkItemRedisWorker][UnsafeGetInstanceKey] redis: nil, fields:[%s, %d]", "C23D7091B98844659D128773209BBF85", 2342324)
	}

	FlushAll()

	if b.N > 1000 {
		b.Logf("cycle:%d", b.N)
		for name, lg := range loggerMap {
			b.Logf("perf:%s:%-5d", name, lg.persistCounter)
		}
	}
}
