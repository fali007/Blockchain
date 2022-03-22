package main

import (
	"os"
	"fmt"
	"math"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"crypto/sha256"
)

type HashObj struct{
	Hash         []byte  `json:"hash"`
	Obj 		 Object  `json:"obj"`
}

type Object struct{
	Nonce        int64   `json:"nonce"`
	Previous     []byte  `json:"previous"`
	Data         string  `json:"data"`
}

func GetJsonEncoding(o interface{})[]byte{
	json,err:=json.Marshal(o)
	if err!=nil{
		fmt.Println("Error marshalling json :",err)
	}
	return json
}

func addBuffer(j []byte, l int)[]byte{
	diff:=l-len(j)
	buf:=make([]byte,diff)
	j=append(j,buf...)
	return j
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

func getDoc()HashObj{
	file,_:=os.Open("db.json")
	b,_:=ioutil.ReadAll(file)
	var temp HashObj
	json.Unmarshal(b,&temp)
	return temp
}

func checkValid(b []byte)bool{
	// fmt.Println(b[:2])
	if bytes.Compare(b[:3], []byte{0,0,0})==0{
		return true
	}
	return false
}

func main(){
	// m:="hello"
	// file,_:=os.OpenFile("db.json",os.O_APPEND|os.O_WRONLY|os.O_CREATE,0600)
	// var i int64=0
	// var prev []byte
	// for i<10{
	// 	obj:=Object{i,prev,m}
	// 	h:=GetSignature(obj)
	// 	file.Write(GetJsonEncoding(HashObj{h,obj}))
	// 	file.Write([]byte("\n"))
	// 	i++
	// 	prev=h
	// }
	// file.Close()
	file,_:=os.OpenFile("db.json",os.O_APPEND|os.O_WRONLY|os.O_CREATE,0600)
	doc:=getDoc()
	fmt.Printf("%+v\n",doc)
	var i int64=0
	for i < math.MaxInt64{
		doc.Obj.Nonce=i
		h:=GetSignature(doc.Obj)
		if checkValid(h){
			fmt.Println("found")
			file.Write(GetJsonEncoding(HashObj{h,doc.Obj}))
			file.Write([]byte("\n"))
			break
		}
		if i%10000==0{
			fmt.Println(i)
		}
		i++
	}
	file.Close()

}