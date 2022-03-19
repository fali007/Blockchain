package service

import (
	"math/rand"
	"time"
	"fmt"
	"unsafe"
	"net/http"
	"encoding/json"
	"serviceA.com/types"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    
	letterIdxMask = 1<<letterIdxBits - 1 
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandomString(n int) string {
    b := make([]byte, n)
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return *(*string)(unsafe.Pointer(&b))
}

func getItems(n int)[]string{
	var a []string
	for i:=0;i<n;i++{
		a=append(a,RandomString(5))
	}
	return a;
}

func getRandomInt(n int)int{
	return rand.Intn(n)
}

func getByte(a interface{}) []byte{
	return []byte(fmt.Sprintf("%+v",a))
}

func PlaceOrder(w http.ResponseWriter, r *http.Request, f bool){
	order:= types.Order{RandomString(4),getItems(getRandomInt(2)),getRandomInt(10)}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(order)
	if f==true{
		fmt.Println("encrypting...")
		cipher:=EncryptWithPublicKey(jsonResp)
		response:=types.EncryptedOrder{cipher}
		resp, _ := json.Marshal(response)
		w.Write(resp)
	}else{
		w.Write(jsonResp)
	}
}