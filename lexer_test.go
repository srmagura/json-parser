package main

import (
	"testing"
)

func TestDoubleNumber(t *testing.T) {
	if DoubleNumber(3) != 6 {
		t.Fatalf("DoubleNumber(3) != 6")
	}
}
