package main

import (
	"os"
	"testing"
)

func TestNewCSVDB(t *testing.T) {
	file, err := os.Open("./engineering_project_businesses.csv")
	if err != nil {
		t.Fatalf("file open error: %#v", err)
	}

	csvdb, err := NewCSVDB(file)
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}
	if csvdb == nil {
		t.Error("nil db")
	}
}
