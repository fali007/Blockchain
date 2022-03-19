package keys

import(
	"crypto/rand"
    "crypto/rsa"
    "encoding/pem"
    "crypto/x509"
    "errors"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

type KeyResponse struct{
    PublicKey   string
    PrivateKey  string
    User        string
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
    privkey, _ := rsa.GenerateKey(rand.Reader, 1024)
    d1 := []byte(ExportRsaPrivateKeyAsPemStr(privkey))
    d2 := []byte(ExportRsaPublicKeyAsPemStr(&privkey.PublicKey))
    ioutil.WriteFile("pem/key.pri", d1, 0644)
    ioutil.WriteFile("pem/key.pub", d2, 0644)
    return privkey, &privkey.PublicKey
}

func GetRsaKeyPair(user string)KeyResponse{
    res,_:=http.Get("http://localhost:5050/getKeys?user="+user)
    var key KeyResponse
    json.NewDecoder(res.Body).Decode(&key)
    return key
}

func GetStoredRsaKeyPair()(*rsa.PrivateKey,*rsa.PublicKey){
    pri,_:= ioutil.ReadFile("pem/key.pri")
    pub,_:= ioutil.ReadFile("pem/key.pub")
    priKey,_:=ParseRsaPrivateKeyFromPemStr(string(pri))
    pubKey,_:=ParseRsaPublicKeyFromPemStr(string(pub))
    return priKey,pubKey
}

func GetStoredPublicKey() string{
    pub,_:= ioutil.ReadFile("pem/key.pub")
    return string(pub)
}

func GetStoredPrivateKey() string{
    pub,_:= ioutil.ReadFile("pem/key.pri")
    return string(pub)
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
    privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
    privkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PRIVATE KEY",
                    Bytes: privkey_bytes,
            },
    )
    return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }
    return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
    pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
    if err != nil {
            return ""
    }
    pubkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PUBLIC KEY",
                    Bytes: pubkey_bytes,
            },
    )
    return string(pubkey_pem)
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
    block, _ := pem.Decode([]byte(pubPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    switch pub := pub.(type) {
    case *rsa.PublicKey:
            return pub, nil
    default:
            break // fall through
    }
    return nil, errors.New("Key type is not RSA")
}