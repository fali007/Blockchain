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
}

type Genesis struct{
	GenesisTime		time.Time			`json:"genesis-time"`
	ChainId 		string				`json:"chain-id"`
	Balances 		map[Account]uint	`json:"balances"`
}

type State struct{
	TxMemPool       []Tx
	Balances        map[Account]uint
}