package uid

import (
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	minUint = 1000000000
	maxUint = 9999999999
)

var (
	last   uint64 = minUint
	prefix        = strconv.FormatUint(uint64(10000+os.Getpid())%65535, 16) +
		strconv.FormatUint(uint64(10000+os.Getppid())%65535, 16)
)

func NextUint() uint64 {
	id := atomic.AddUint64(&last, 1)
	if id > maxUint {
		atomic.StoreUint64(&last, minUint)
		return 1
	}

	return id
}

func NextLong() string {
	return prefix +
		strconv.FormatInt(time.Now().Unix(), 16) +
		strconv.FormatUint(NextUint(), 16)
}
