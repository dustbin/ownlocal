package main

import(
	"encoding/json"
	"net/http"
	"os"
	"log"
	"strconv"
	"strings"
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

	http.HandleFunc("/business/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.EscapedPath(),"/")
		id,err := strconv.Atoi(s[2])
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: non numeric id"))
			return
		}
		b,err := businessDB.GetBusiness(id)
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: bad id"))
			return
		}
		j,err := json.Marshal(b)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	http.HandleFunc("/businesses/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.EscapedPath(),"/")
		page,size := 0,50
		if s[2]!="" {
			page,err = strconv.Atoi(s[2])
			if err!=nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 error: non numeric page number"))
				return
			}

		}
		if len(s)>3 && s[3]!="" {
			size,err = strconv.Atoi(s[3])
			if err!=nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 error: non numberic page size"))
				return
			}
		}

		b,err := businessDB.GetPage(page,size)
		if err!=nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: bad page number"))
			return
		}

		j,err := json.Marshal(b)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 error"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
