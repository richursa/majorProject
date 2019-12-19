package cryptorsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func WriteRsaPrivateKeyAsPem(privkey *rsa.PrivateKey, fileName string) {
	privkey_pem_str := exportRsaPrivateKeyAsPemStr(privkey)
	err := ioutil.WriteFile(fileName, []byte(privkey_pem_str), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GetRsaPrivateKeyFromPem(fileName string) *rsa.PrivateKey {
	privkey_pem_str, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	pvtkey, err := parseRsaPrivateKeyFromPemStr(string(privkey_pem_str))
	if err != nil {
		fmt.Println(err)
	}
	return pvtkey
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}
func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
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
