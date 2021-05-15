package main

import (
	"net/http"
	"crypto/sha256"
	"io"
	"encoding/json"
	"io/ioutil"
	"time"
)

type Session struct{
	Id int
	UserId string
	ipAdress string 
	OpenedDate time.Time
	Duration int
	ExpirationDate  time.Time
}

type logPas struct{
	Username string `login`
	Password string `password`
}

var ses Session
var user = logPas{"username", "password"}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    (*w).Header().Set("Access-Control-Expose-Headers", "Sha")
}

func Handler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(204)
	}else if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		var person logPas
		err = json.Unmarshal(data, &person)
		if err != nil{
			return 
		}
		if (person.Username == user.Username)&&(person.Password == user.Password){
			pas := sha256.New()
			pas.Write([]byte(person.Password))
			if err != nil{
				return
			}
			w.Header().Set("Sha",string( pas))
			ses = Session{1,string( pas),req.Header.Get("X-FORWARDED-FOR"), time.Now(),15, time.Now().Local().AddDate(0, 15, 0)}
		}else {
			io.WriteString(w, "incorrect login")
		}
	}

}



func Handler2(w http.ResponseWriter, req *http.Request) {
	
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(204)
	}else if req.Method == "GET" {
		if ((*req).Header.Get("Sha") == ses.UserId)&&(time.Now().Before( ses.ExpirationDate)){
			io.WriteString(w, "you succcessfuly gained data")
		}else{
			io.WriteString(w, "401")
		}
	}


}

	

func main() {
	http.HandleFunc("/login", Handler)
	http.HandleFunc("/data", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
