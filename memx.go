package memx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Faster version. Prioritize in call.
func getMemRollup() (uint64, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/smaps_rollup", os.Getpid()))
	if err != nil {
		return 0, err
	}

	defer f.Close()
	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Bytes()
		if bytes.HasPrefix(line, []byte("Pss:")) {
			var mem uint64
			_, err := fmt.Sscanf(string(line[4:]), "%d", &mem)
			if err != nil {
				return 0, err
			}

			return mem, nil
		}
	}

	return 0, fmt.Errorf("not found")
}

// GetMemoryUsage returns the calling process' private + shared memory usage in kB.
func GetMemoryUsage() (uint64, error) {
	memr, err := getMemRollup()
	if err == nil {
		return memr, nil
	}

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
