package main

import(
	"crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "os"
)

type Person struct{
	Name string
	Age int
}

func getByte(a interface{}) []byte{
	return []byte(fmt.Sprintf("%+v",a))
}

func getKey() *rsa.PrivateKey{
	privateKey,err:=rsa.GenerateKey(rand.Reader,1024)
	if err!=nil{
		fmt.Println("Error generating key ",err.Error)
		os.Exit(1)
	}
	// fmt.Printf("\n\n%+v\n\n",privateKey)
	return privateKey;
}

func main(){
	felix:=getKey()
	hari:=getKey()

	person:=Person{Name:"felix",Age:12}
	
	message:=getByte(person)
	label:=getByte("")

	h:=sha256.New()

	cipherText,err:=rsa.EncryptOAEP(h,rand.Reader,&felix.PublicKey,message,label)
	if err!=nil{
		fmt.Println("Error generating key ",err.Error)
		os.Exit(1)
	}

	fmt.Printf("Encrypted message to %x\n",cipherText)

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader,hari,newhash,hashed,&opts)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("PSS Signature : %x\n", signature)

	plainText, err := rsa.DecryptOAEP(h,rand.Reader,felix,cipherText,label)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", cipherText, plainText)

	err = rsa.VerifyPSS(&hari.PublicKey,newhash,hashed,signature,&opts)
	if err != nil {
	    fmt.Println("Who are U? Verify Signature failed")
	    os.Exit(1)
	} else {
	    fmt.Println("Verify Signature successful")
	}
}