package service

import (
	"blockchain/types"
)

var transactionChan = make(chan types.Tx)

func GetChannel()*chan types.Tx{
	return &transactionChan
}

func CloseChannel()bool{
	close(*GetChannel())
	return true
}

func StartChannel()bool{
	c:=GetChannel()
	go Process(c)
	return true
}

func Process(c *chan types.Tx){
	for doc:=range *c{
		validateTransation(doc)
	}
}