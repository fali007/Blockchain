package service

import(
	"fmt"
	"os"
    "crypto/rand"
	"crypto/rsa"
    "crypto/sha256"
    "serviceA.com/keys"
)

func EncryptWithPublicKey(m []byte)[]byte{
	fmt.Printf("message to encrypt %x\n",m)
	_,key:=keys.GetStoredRsaKeyPair()

	cipherText,err:=rsa.EncryptOAEP(sha256.New(),rand.Reader,key,m,getByte(""))
	if err!=nil{
		fmt.Println("Error generating key ",err)
		os.Exit(1)
	}

	fmt.Printf("Encrypted message to %x\n",cipherText)
	return cipherText
}