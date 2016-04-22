package main

import(
	"encoding/json"
	"fmt"
	"os"
	"log"
	"github.com/dustbin/ownlocal/business"
)

func main(){
	file,err := os.Open("./engineering_project_businesses.csv")
	if(err!=nil){
		log.Fatal(err)
	}
	defer file.Close()

	businessDB,err := business.NewBusinessDB(file)
	if(err!=nil){
		log.Fatal(err)
	}
	b,err := businessDB.GetBusiness(0)
	fmt.Println(b)
	j,err := json.Marshal(b)
	fmt.Println(string(j))
}
