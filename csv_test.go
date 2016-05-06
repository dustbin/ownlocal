package main

import (
	"os"
	"testing"
)

func TestNewCSVDB(t *testing.T) {
	file, err := os.Open("./engineering_project_businesses.csv")
	if err != nil {
		t.Fatal(err)
	}

	csvdb, err := NewCSVDB(file)
	if err != nil {
		t.Error(err)
	}
	if csvdb == nil {
		t.Error("nil db")
	}
}
