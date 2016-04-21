package main

import(
	"github.com/dustbin/ownlocal/csv"
	"fmt"
	"os"
	"log"
)

func main(){
	file,err := os.Open("./engineering_project_businesses.csv")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()

	csvdb,err := csv.NewCSVDB(file)
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println(csvdb)
}
