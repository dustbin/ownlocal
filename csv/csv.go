package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

//holds header row and data rows obtained from a CSV
type CSVDB struct {
	headers []string
	rows    [][]string
}

//creates a new CSVDB object from a io.Reader, such as a File obtained by os.Open
func NewCSVDB(reader io.Reader) (*CSVDB, error) {
	csvdb := CSVDB{}

	csvr := csv.NewReader(reader)
	headers, err := csvr.Read()
	if err != nil {
		return &csvdb, err
	}
	csvdb.headers = headers
	records, err := csvr.ReadAll()
	if err != nil {
		return &csvdb, err
	}
	csvdb.rows = records

	return &csvdb, nil
}

//returns the number of rows
func (c *CSVDB) GetSize() int {
	return len(c.rows)
}

//returns a specific row
func (c *CSVDB) GetRow(i int) []string {
	return c.rows[i]
}

//keeps the print of the CSVDB manageable
func (c *CSVDB) String() string {
	return fmt.Sprintf("%v\n%d rows", c.headers, len(c.rows))
}
