package main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	// flag.Parse()
	os.Exit(m.Run())
}
