package types

import(
	"time"
)

type Account string

type Tx struct{
	From       		Account     		`json:"from"`
	To         		Account     		`json:"to"`
	Value      		uint        		`json:"value"`
	Data       		string      		`json:"data"`
	Time 			time.Time 			`json:"time"`
}

type TxDoc struct{
	Nonce			int64 				`json:"nonce"`
	Previous 		[]byte 				`json:"previous"`
	Transaction 	Tx 					`json:"transaction"`
}

type HashObj struct{
	Signature		[]byte				`json:"sign"`
	Document 		TxDoc				`json:"doc"`
}

type Genesis struct{
	GenesisTime		time.Time			`json:"genesis-time"`
	ChainId 		string				`json:"chain-id"`
	Balances 		map[Account]uint	`json:"balances"`
}

type State struct{
	TxMemPool       []HashObj
	Balances        map[Account]uint
	LastHash		[]byte
}