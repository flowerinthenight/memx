package memx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// GetMemoryUsage returns the calling process' internal + shared
// memory usage. Linux-specific as it reads the /proc/... files.
func GetMemoryUsage() (uint64, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/smaps", os.Getpid()))
	if err != nil {
		return 0, err
	}

	defer f.Close()
	ret := uint64(0)
	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Bytes()
		if bytes.HasPrefix(line, []byte("Pss:")) {
			var size uint64
			_, err := fmt.Sscanf(string(line[4:]), "%d", &size)
			if err != nil {
				return 0, err
			}

			ret += size
		}
	}

	if r.Err() != nil {
		return 0, err
	}

	return ret, nil
}
