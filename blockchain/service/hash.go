package service

import(
	"fmt"
	"math"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"blockchain/types"
)

func addBuffer(j []byte, l int)[]byte{
	diff:=l-len(j)
	buf:=make([]byte,diff)
	j=append(j,buf...)
	return j
}

func GetValidSignatureBlock(h *types.HashObj)*types.HashObj{
	var i int64=0
	for i < math.MaxInt64{
		(*h).Document.Nonce=i
		s:=GetSignature((*h).Document)
		if checkValid(s){
			(*h).Signature=s
			fmt.Println("found")
			break
		}
		i++
	}
	return h
}

func GetSignature(i interface{})[]byte{
	sign_len:=64
	j,_:=json.Marshal(i)
	if len(j)<sign_len{
		j=addBuffer(j,sign_len)
	}
	h:=sha256.New()
	k:=sign_len
	l:=0
	for k<len(j){
		h.Write(j[l:k])
		l=k
		k+=sign_len
	}
	h.Write(j[l:])
	return h.Sum(nil)
}

func checkValid(b []byte)bool{
	if bytes.Compare(b[:3], []byte{0,0,0})==0{
		return true
	}
	return false
}