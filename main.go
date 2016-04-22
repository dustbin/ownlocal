package main

import(
	"encoding/json"
	"net/http"
	"os"
	"log"
	"strconv"
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

	http.HandleFunc("/business", func(w http.ResponseWriter, r *http.Request) {
		idVal := r.FormValue("id")
		id,err := strconv.Atoi(idVal)
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b,err := businessDB.GetBusiness(id)
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		j,err := json.Marshal(b)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	http.HandleFunc("/businesses", func(w http.ResponseWriter, r *http.Request) {
		page,size := 0,50
		pageVal := r.FormValue("page")
		if pageVal!="" {
			page,err = strconv.Atoi(pageVal)
			if err!=nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		sizeVal := r.FormValue("size")
		if sizeVal!="" {
			size,err = strconv.Atoi(sizeVal)
			if err!=nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		b,err := businessDB.GetPage(page,size)
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		j,err := json.Marshal(b)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
