package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"

	"../cryptorsa"
)

func main() {
	fmt.Println("Enter The NodeID to generate private key")
	var NodeID string
	fmt.Scanln(&NodeID)
	var bits int
	fmt.Println("Enter the number of bits of rsa key")
	fmt.Scanln(&bits)
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatalln(err)
	}
	cryptorsa.WriteRsaPrivateKeyAsPem(privateKey, NodeID+".pem")
}
