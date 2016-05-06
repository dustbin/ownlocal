package main

import (
	"testing"
)

func TestLoadCSV(t *testing.T) {
	businessDB,err := loadCSV()
	if err!=nil {
		t.Error(err)
	}
	if businessDB==nil {
		t.Error("nil db")
	}
}
