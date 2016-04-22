package csv

import(
	"fmt"
	"encoding/csv"
	"io"
)

type CSVDB struct {
	headers []string
	rows [][]string
}

func NewCSVDB(reader io.Reader) (*CSVDB,error) {
	csvdb := CSVDB{}

	csvr := csv.NewReader(reader)
	headers,err := csvr.Read()
	if(err!=nil){
		return &csvdb,err
	}
	csvdb.headers = headers
	records,err := csvr.ReadAll()
	if(err!=nil){
		return &csvdb,err
	}
	csvdb.rows = records

	return &csvdb,nil
}

func (c *CSVDB) GetSize() int{
	return len(c.rows)
}

func (c *CSVDB) GetRow(i int) []string{
	return c.rows[i]
}

func (c *CSVDB) String() string{
	return fmt.Sprintf("%v\n%d rows",c.headers,len(c.rows))
}
