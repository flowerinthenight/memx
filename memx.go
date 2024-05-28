package memx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// GetMemoryUsage returns the calling process' internal + shared
// memory usage in kB. Linux-specific as it reads /proc/.
// Ref: https://en.wikipedia.org/wiki/Proportional_set_size
func GetMemoryUsage() (uint64, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/smaps", os.Getpid()))
	if err != nil {
		return 0, err
	}

	defer f.Close()
	var mem uint64
	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Bytes()
		if !bytes.HasPrefix(line, []byte("Pss:")) {
			continue
		}

		var size uint64
		var unit string
		_, err := fmt.Sscanf(string(line[4:]), "%d %s", &size, &unit)
		if err != nil {
			return 0, err
		}

		if strings.ToLower(unit) == "kb" {
			mem += size
		}
	}

	if r.Err() != nil {
		return 0, err
	}

	return mem, nil
}
