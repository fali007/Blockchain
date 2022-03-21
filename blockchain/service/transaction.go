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

	if doc.Data=="t"{
		writeToDb(GetJsonEncoding(types.TxDoc{doc,10000,GetSignature(doc),state.LastHash}))
		return
	}
	
	if state.Balances[doc.From]>=doc.Value{
		writeToDb(GetJsonEncoding(types.TxDoc{doc,10000,GetSignature(doc),state.LastHash}))
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

func getDb(b *map[types.Account]uint, h *[]byte)*[]types.Tx{
	tx_document,err:=os.Open("db.json")
	defer tx_document.Close()
	
	if err!=nil{
		fmt.Println("Error opening info file", err)
	}

	var tx []types.Tx
	var hash []byte

	scanner := bufio.NewScanner(tx_document)
	for scanner.Scan(){
		var temp types.TxDoc
		json.Unmarshal(scanner.Bytes(),&temp)
		if temp.Transaction.Data=="f"{
			(*b)[temp.Transaction.From]-=temp.Transaction.Value
			(*b)[temp.Transaction.To]+=temp.Transaction.Value
		}else{
			(*b)[temp.Transaction.To]+=temp.Transaction.Value
		}
		hash=temp.Signature
        tx=append(tx,temp.Transaction)
    }
    *h=hash
    return &tx
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