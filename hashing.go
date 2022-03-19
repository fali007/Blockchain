package main

import(
	"fmt"
	"crypto/sha256"
)

func getHash(s string)string{
	h:=sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x",h.Sum(nil))
}

func main(){
	fmt.Println(getHash("hello world"))
}