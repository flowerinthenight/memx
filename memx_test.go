package memx

import "testing"

func Test_GetMemoryUsage(t *testing.T) {
	v, err := GetMemoryUsage()
	if err != nil {
		t.Fatalf("GetMemoryUsage failed: expected nil, got %v.", err)
	}

	if v == 0 {
		t.Fatalf("Expected non-zero, got %v.", v)
	}
}
