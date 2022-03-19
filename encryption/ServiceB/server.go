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

func getOrder(w http.ResponseWriter, r *http.Request){
	service.GetOrder(w,r)
}

func getEncrytedOrder(w http.ResponseWriter, r *http.Request){
	service.GetEncryptedOrder(w,r)
}

func main(){
	// keys.GenerateRsaKeyPair()
	// fmt.Printf("%+v\n",keys.GetRsaKeyPair("user"))
	fmt.Println("Generated keys. Starting server B")
	r:=mux.NewRouter()
	r.HandleFunc("/getPublicKey", PublicKey).Methods("GET")
	r.HandleFunc("/getorder", getOrder).Methods("GET")
	r.HandleFunc("/getencryptedorder", getEncrytedOrder).Methods("GET")
	http.ListenAndServe(":8081", r)
}