package service

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"blockchain/types"
)

func writeToDb(s []byte)bool{
	status:=true
	file,err:=os.OpenFile("db.json",os.O_APPEND|os.O_WRONLY|os.O_CREATE,0600)
	if err!=nil{
		status=false
	}
	_,err=file.Write(s)
	_,err=file.Write([]byte("\n"))
	if err!=nil{
		status=false
	}
	return status
}

func validateTransation(doc types.Tx){
	state:=LoadState()
	var nonce int64=1

	if doc.Data=="t"{
		document:=types.TxDoc{nonce,state.LastHash,doc}
		hashObj:=types.HashObj{GetSignature(document),document}
		adress:=GetValidSignatureBlock(&hashObj)
		writeToDb(GetJsonEncoding(*adress))
		return
	}
	
	if state.Balances[doc.From]>=doc.Value{
		document:=types.TxDoc{nonce,state.LastHash,doc}
		hashObj:=types.HashObj{GetSignature(document),document}
		adress:=GetValidSignatureBlock(&hashObj)
		writeToDb(GetJsonEncoding(*adress))
		return
	}else{
		fmt.Println("Invalid or Insufficient Transaction")
	}
}

func Transaction(f,t,d,v string)bool{
	c:=GetChannel()
	value,err:=strconv.ParseUint(v,10,64)
	if err!=nil{
		fmt.Println("Value Not a number :",err)
	}
	log:=types.Tx{types.Account(f),types.Account(t),uint(value),d,time.Now()}
	*c <- log
	if err!=nil{
		return false
	}
	return true
}

func getGenesis()*types.Genesis{
	document,err:=ioutil.ReadFile("genesis.json")
	if err!=nil{
		fmt.Println("Error opening info file", err)
	}
	var genesis types.Genesis
	json.Unmarshal(document,&genesis)
	return &genesis
}

func getDb(b *map[types.Account]uint, h *[]byte)*[]types.HashObj{
	tx_document,err:=os.Open("db.json")
	defer tx_document.Close()
	
	if err!=nil{
		fmt.Println("Error opening info file", err)
	}

	var tx []types.HashObj
	var hash []byte

	scanner := bufio.NewScanner(tx_document)
	for scanner.Scan(){
		var temp types.HashObj
		json.Unmarshal(scanner.Bytes(),&temp)
		if temp.Document.Transaction.Data=="f"{
			(*b)[temp.Document.Transaction.From]-=temp.Document.Transaction.Value
			(*b)[temp.Document.Transaction.To]+=temp.Document.Transaction.Value
		}else{
			(*b)[temp.Document.Transaction.To]+=temp.Document.Transaction.Value
		}
		hash=temp.Signature
        tx=append(tx,temp)
    }
    *h=hash
    return &tx
}

func isEqual(a,b []byte)bool{
	if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func ValidateStateWithResponse()[]byte{
	state:=LoadState()
	root:=true
	var previous []byte=nil
	for _,v:=range state.TxMemPool{
		if !checkValid(v.Signature){
			return GetJsonEncoding(types.ValidateResponse{"failed",string(GetJsonEncoding(v))})
		}
		if root{
			if !isEqual(v.Signature,GetSignature(v.Document)){
				return GetJsonEncoding(types.ValidateResponse{"failed",string(GetJsonEncoding(v))})
			}
		}else{
			if !isEqual(v.Signature,GetSignature(v.Document))||!isEqual(v.Document.Previous,previous){
				return GetJsonEncoding(types.ValidateResponse{"failed",string(GetJsonEncoding(v))})
			}
		}
		previous=v.Signature
		root=false
	}
	return GetJsonEncoding(types.ValidateResponse{"success","nil"})
}

func ValidateState()bool{
	state:=LoadState()
	root:=true
	var previous []byte=nil
	for _,v:=range state.TxMemPool{
		if !checkValid(v.Signature){
			return false
		}
		if root{
			if !isEqual(v.Signature,GetSignature(v.Document)){
				return false
			}
		}else{
			if !isEqual(v.Signature,GetSignature(v.Document))||!isEqual(v.Document.Previous,previous){
				return false
			}
		}
		previous=v.Signature
		root=false
	}
	return true
}

func LoadState()*types.State{
	genesis:=getGenesis()
	
	var state types.State
	state.Balances = make(map[types.Account]uint)

	for k,v:=range genesis.Balances{
		state.Balances[k]=v
	}
	state.TxMemPool=*getDb(&state.Balances, &state.LastHash)
	return &state
}