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
	Transaction 	Tx 					`json:"transaction"`
	Nonce			uint 				`json:"nonce"`
	Signature 		[]byte				`json:"signature"`
	Previous 		[]byte 				`json:"previous"`
}

type Genesis struct{
	GenesisTime		time.Time			`json:"genesis-time"`
	ChainId 		string				`json:"chain-id"`
	Balances 		map[Account]uint	`json:"balances"`
}

type State struct{
	TxMemPool       []TxDoc
	Balances        map[Account]uint
	LastHash		[]byte
}