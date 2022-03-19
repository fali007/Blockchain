package types

type KeyResponse struct{
	Key string
	Message string
}

type Order struct{
	OrderId string
	Items []string
	Cost int
}

type EncryptedOrder struct{
	Order []byte
}