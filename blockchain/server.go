package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"blockchain/service"
)

func Index(w http.ResponseWriter, r *http.Request){
	service.IndexController(w,r)
}

func Genesis(w http.ResponseWriter, r *http.Request){
	service.GenesisController(w,r)
}

func Transaction(w http.ResponseWriter, r *http.Request){
	service.TransactionController(w,r)
}

func Close(w http.ResponseWriter, r *http.Request){
	service.CloseTransaction(w,r)
}

func main(){
	if f:=service.StartChannel(); f==true{
		fmt.Println("State loaded and Server started")
	}
	r:=mux.NewRouter()
	r.HandleFunc("/index",Index).Methods("GET")
	r.HandleFunc("/genesis",Genesis).Methods("GET")
	r.HandleFunc("/transaction",Transaction).Methods("GET")
	r.HandleFunc("/close",Close).Methods("GET")
	http.ListenAndServe(":8080",r)
}