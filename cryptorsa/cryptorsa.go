package cryptorsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func WriteRsaPrivateKeyAsPem(privkey *rsa.PrivateKey, fileName string) {
	privkey_pem_str := exportRsaPrivateKeyAsPemStr(privkey)
	err := ioutil.WriteFile(fileName, []byte(privkey_pem_str), 0777)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetRsaPrivateKeyFromPem(fileName string) *rsa.PrivateKey {
	privkey_pem_str, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	pvtkey, err := parseRsaPrivateKeyFromPemStr(string(privkey_pem_str))
	if err != nil {
		log.Fatalln(err)
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
func WriteRsaPublicKeyAsPem(publickey *rsa.PublicKey, fileName string) {
	publickey_pem_str, err := exportRsaPublicKeyAsPemStr(publickey)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(fileName, []byte(publickey_pem_str), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GetRsaPublicKeyFromPem(fileName string) *rsa.PublicKey {
	publickey_pem_str, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	publickey, err := ParseRsaPublicKeyFromPemStr(string(publickey_pem_str))
	if err != nil {
		log.Fatalln(err)
	}
	return publickey
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
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
