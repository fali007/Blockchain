package service

import(
	"fmt"
	"crypto/sha256"
	"encoding/json"
)

func GetSignature(i interface{})[]byte{
	txJson,err:=json.Marshal(i)
	if err!=nil{
		fmt.Println("Error marshalling object", err)
	}
	h:=sha256.New()
	return h.Sum(txJson)
}