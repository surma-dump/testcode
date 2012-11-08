package main

import (
	"testing"
)

func TestNumber(t *testing.T) {
	if Number() != 4 {
		t.Fatalf("AAAH!")
	}
}
