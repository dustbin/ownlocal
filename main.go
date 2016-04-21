package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"log"
)
type CSVDB struct {
	headers []string
	rows [][]string
}

func main(){
	file,err := os.Open("./engineering_project_businesses.csv")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()
	csvdb := CSVDB{}
	csvr := csv.NewReader(file)
	headers,err := csvr.Read()
	if(err!=nil){
		log.Fatal(err)
	}
	csvdb.headers = headers
	records,err := csvr.ReadAll()
	if(err!=nil){
		log.Fatal(err)
	}
	csvdb.rows = records

	fmt.println(csvdb)
}
