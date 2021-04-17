package main

import (
	"net/http"
	//"io"
//	"io/ioutil"
//	"fmt"
)
var id []int

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Handler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	/*pas := sha256.New()
	pas.Write([]byte(//password))
	if (//login == "login")&&(//pas.Sum(nil) == ae.....){
		//разрешить доступ 
	}else{
		io.WriteString(w, "unsuccessful login")
	}*/
	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	/*if (correctID(idn)){
		io.WriteString(w, "you succcessfuly gained data")
	}else{
		io.WriteString(w, "you unsucccessfuly gained data")
	}*/


}

func correctID(idn int) bool{
	for i := 0; i< len(id); i++{
		if idn == id[i]{
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
