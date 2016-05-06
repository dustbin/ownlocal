package main

import (
	"encoding/csv"
	"io"
)

//holds header row and data rows obtained from a CSV
type CSVDB struct {
	Headers []string
	Rows    [][]string
}

//creates a new CSVDB object from a io.Reader, such as a File obtained by os.Open
func NewCSVDB(reader io.Reader) (*CSVDB, error) {
	csvdb := CSVDB{}

	csvr := csv.NewReader(reader)
	headers, err := csvr.Read()
	if err != nil {
		return &csvdb, err
	}
	csvdb.Headers = headers
	records, err := csvr.ReadAll()
	if err != nil {
		return &csvdb, err
	}
	csvdb.Rows = records

	return &csvdb, nil
}
