package main

import (
	"net/http"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"fmt"
)

type Session struct{
	Id int
	UserId int
	idAdress string 
	OpenedDate string
	Duration string
	ExpirationDate  string
}

type logPas struct{
	Login string `login`
	Password string `password`
}

var session []Session 
var idn int

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Handler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	pas := sha256.New()
	data, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	var person logPas
	err = json.Unmarshal(data, &person)
	if err != nil{
		return err
	}
	pas.Write([]byte(person.Password))
	if (person.Login == "login")&&(pas.Sum(nil) == "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"){
		session[len(session)] = {1, 1, "id1029e", "17.04.2021", "3", "18.04.2021"}
		idn = 1
	}else{
		io.WriteString(w, "unsuccessful login")
	}
	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	if (correctID(idn)){
		io.WriteString(w, "you succcessfuly gained data")
	}else{
		io.WriteString(w, "you unsucccessfuly gained data")
	}


}

func correctID(idn int) bool{
	for i := 0; i< len(session); i++{
		if idn == session[i].Id{
			return true 
		}
	}
	return false
}
	

func main() {
	http.HandleFunc("/login", Handler)
	http.HandleFunc("/data", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
