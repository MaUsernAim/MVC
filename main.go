package main

import "net/http"

func main(){

	//firstHANDLE 
	http.HandleFunc("/",mainPage)

	//secondary handler
	http.HandleFunc("data/dataFile",sendLocalData)
	http.ListenAndServe(":8002,nil")
}