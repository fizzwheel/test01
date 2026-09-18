package demo01

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", Add(1, 2))
	}
}

func Add(a, b int) int {
	return a + b
}
