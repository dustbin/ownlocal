package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	businessDB, err := loadCSV()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	//handles requests to business
	router.HandleFunc("/business/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.EscapedPath(), "/")
		id, err := strconv.Atoi(s[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: non numeric id"))
			return
		}
		b, err := businessDB.GetBusiness(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: bad id"))
			return
		}
		j, err := json.Marshal(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	//handles requests to businesses
	router.HandleFunc("/businesses/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.EscapedPath(), "/")
		page, size := 0, 50
		if s[2] != "" {
			page, err = strconv.Atoi(s[2])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 error: non numeric page number"))
				return
			}

		}
		if len(s) > 3 && s[3] != "" {
			size, err = strconv.Atoi(s[3])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 error: non numberic page size"))
				return
			}
		}

		b, err := businessDB.GetPage(page, size)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 error: bad page number"))
			return
		}

		j, err := json.Marshal(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	//catch all for other requests
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 error"))
	})

	n := negroni.New()
	n.Use(negroni.HandlerFunc(MyMiddleware))
	n.UseHandler(router)
	n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, password, ok := r.BasicAuth()
	rw.Header().Add("WWW-Authenticate","Basic")
	if !ok {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("401 error: no auth provided"))
		return
	}
	if password != "token" {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("401 error: bad token"))
		return
	}
	next(rw, r)
}

func loadCSV() (*BusinessDB, error) {
	//open csv file
	file, err := os.Open("./engineering_project_businesses.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//creates BusinessDB object
	businessDB, err := NewBusinessDB(file)
	return businessDB, err
}
