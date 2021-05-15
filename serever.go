package main

import (
	"net/http"
	"crypto/sha256"
	"io"
	"encoding/json"
//	"encoding/hex"
	"io/ioutil"
	"time"
	"fmt"
)

type Session struct{
	Id int
	UserId int
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
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionId")
    (*w).Header().Set("Access-Control-Expose-Headers", "SessionId")
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
		pas := sha256.Sum256([]byte(user.Password))
		pas2 := sha256.Sum256([]byte(person.Password))
		if (person.Username == user.Username)&&(pas == pas2){
			pas := sha256.New()
			pas.Write([]byte(person.Password))
			if err != nil{
				return
			}
			ses = Session{1,1,req.Header.Get("X-FORWARDED-FOR"), time.Now(),15, time.Now().Local().AddDate(0, 15, 0)}
			w.Header().Set("SessionId",string( ses.Id))
			io.WriteString(w, "successfully login")
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
		fmt.Println((*req).Header.Get("SessionId"))
		if ((*req).Header.Get("SessionId") == string(ses.Id))&&(time.Now().Before( ses.ExpirationDate)){
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
