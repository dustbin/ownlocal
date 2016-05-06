package main

import (
	"testing"
)

func TestLoadCSV(t *testing.T) {
	businessDB, err := loadCSV()
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}
	if businessDB == nil {
		t.Error("nil db")
	}
}
