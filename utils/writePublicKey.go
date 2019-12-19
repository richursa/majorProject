package main

import (
	"fmt"

	"../cryptorsa"
)

func main() {
	fmt.Println("Enter the name/path of the PrivateKey")
	var filename string
	fmt.Scanln(&filename)
	pvtKey := cryptorsa.GetRsaPrivateKeyFromPem(filename)
	cryptorsa.WriteRsaPublicKeyAsPem(&pvtKey.PublicKey, filename+".pub")
}
