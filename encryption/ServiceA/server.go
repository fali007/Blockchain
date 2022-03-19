package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "serviceA.com/keys"
    "serviceA.com/types"
    "serviceA.com/service"
)

func PublicKey(w http.ResponseWriter, r *http.Request){
	pub:=keys.GetStoredPublicKey()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp:=types.KeyResponse{Key:pub,Message:"serviceA PublicKey"}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func Order(w http.ResponseWriter, r *http.Request){
	service.PlaceOrder(w,r,false)
}

func EncrytedOrder(w http.ResponseWriter, r *http.Request){
	service.PlaceOrder(w,r,true)
}

func main(){
	// keys.GenerateRsaKeyPair()
	// fmt.Printf("%+v\n",keys.GetRsaKeyPair("user"))
	fmt.Println("Generated keys. Starting server A")
	r:=mux.NewRouter()
	r.HandleFunc("/getPublicKey", PublicKey).Methods("GET")
	r.HandleFunc("/order", Order).Methods("GET")
	r.HandleFunc("/encryptedorder", EncrytedOrder).Methods("GET")
	http.ListenAndServe(":8080", r)
}