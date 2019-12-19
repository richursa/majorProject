package verification

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"

	"../blockchain"
	"../cryptorsa"
)

// implement additional verifications like if prev block matches
func VerifyBlock(block blockchain.Block) bool {
	publickey := cryptorsa.GetRsaPublicKeyFromPem("/app/nodeinfo/publickeys/" + block.NodeID + ".pem.pub")
	data := []byte(block.Data)
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publickey, crypto.SHA256, hashed[:], block.Signature)
	if err != nil {
		return false
	} else {
		return true
	}
}
