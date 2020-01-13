package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"reflect"

	"../cryptorsa"
)

func main() {
	pvtKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	cryptorsa.WriteRsaPrivateKeyAsPem(pvtKey, "richu.pem")
	message := []byte("richu has completed btech")
	hashed := sha256.Sum256(message)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, pvtKey, crypto.SHA256, hashed[:])
	fmt.Println(reflect.TypeOf(signature))
	pvtKey = cryptorsa.GetRsaPrivateKeyFromPem("richu.pem")
	err := rsa.VerifyPKCS1v15(&pvtKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("signature not valid")
	} else {
		fmt.Println("signature valid")
	}
}
