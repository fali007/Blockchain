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
		// go validateTransation(doc) -- running multi threaded gives an error - previous hash will be mismatch
		validateTransation(doc)
	}
}