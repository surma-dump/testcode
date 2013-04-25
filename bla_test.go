package main

import (
	"testing"

	"github.com/surma-dump/gouuid"
)

func TestNumber(t *testing.T) {
	_ = gouuid.New()
	if Number() != 4 {
		t.Fatalf("AAAH!")
	}
}
