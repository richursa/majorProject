package verification

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"

	"../blockchain"
	"../cryptorsa"
	"../db"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyBlock(block blockchain.Block) bool {
	mongoClient := db.GetClient()
	prevBlock := db.GetBlockFromDB(mongoClient, bson.M{"blockid": db.GetCount(mongoClient)})
	toBeHashed := blockchain.IntToStr(block.BlockID) + blockchain.IntToStr(block.Time) + block.Data + string(block.Signature) + block.NodeID
	if block.BlockID != prevBlock.BlockID+1 {
		return false
	} else if block.Time < prevBlock.Time {
		return false
	} else if VerifySignature(block) == false {
		return false
	} else if block.Prev != prevBlock.Hash {
		return false
	} else if blockchain.CalcHash(blockchain.IntToStr(block.Nonce)+toBeHashed) != block.Hash {
		return false
	} else {
		return true
	}
}

func VerifySignature(block blockchain.Block) bool {
	publickey := cryptorsa.GetRsaPublicKeyFromPem("/app/nodeinfo/publickeys/" + block.NodeID + ".pem.pub")
	data := []byte(block.Data)
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publickey, crypto.SHA256, hashed[:], block.Signature)
	if err != nil {
		return false
	}
	return true
}
