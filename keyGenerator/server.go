package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "crypto/rand"
    "crypto/rsa"
    "encoding/pem"
    "crypto/x509"
)

type KeyResponse struct{
	PublicKey   string
	PrivateKey  string
	User        string
}

func generateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
    privkey, _ := rsa.GenerateKey(rand.Reader, 1024)
    return privkey, &privkey.PublicKey
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
    privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
    privkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PRIVATE KEY",
                    Bytes: privkey_bytes,
            },
    )
    return string(privkey_pem)
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
    pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
    if err != nil {
            return ""
    }
    pubkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PUBLIC KEY",
                    Bytes: pubkey_bytes,
            },
    )
    return string(pubkey_pem)
}

func GenerateKey(w http.ResponseWriter, r *http.Request){
	user, ok := r.URL.Query()["user"]
	if !ok || len(user[0]) < 1 {
        fmt.Println("Username is missing")
    }
	key,_:=generateRsaKeyPair()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp:=KeyResponse{PublicKey:exportRsaPublicKeyAsPemStr(&key.PublicKey),PrivateKey:exportRsaPrivateKeyAsPemStr(key),User:user[0]}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func main(){
	fmt.Println("Key Generator started")
	r:=mux.NewRouter()
	r.HandleFunc("/getKeys", GenerateKey).Methods("GET")
	http.ListenAndServe(":5050", r)
}