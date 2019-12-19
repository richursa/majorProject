package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"../nodeinfo"
)

//Block : basic block of the blokchain
type Block struct {
	BlockID   int64   `json:"blockID"`   //stores block number
	Time      int64   `json:"time"`      //to store time
	Data      string  `json:"data"`      //transactions/data which is to be stored in a block
	Signature []uint8 `json:"signature"` //signature of the data
	NodeID    string  `json:"nodeID"`    //id of the node
	Prev      string  `json:"prev"`      //hash of the previous block
	Hash      string  `json:"hash"`      //hash of the current block
	Nonce     int64   `json:"nonce"`     //Nonce of the current block
}

//return sha256 hash of the given string
func calcHash(data string) string {
	hashed := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashed[:])
}

//convert a integer to its string equivalent
func intToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

//compute a hash of given difficulty by incrememting nonce
func computeHashWithProofOfWork(data string, difficulty string) (int64, string) {
	nonce := int64(0)
	for {
		hash := calcHash(intToStr(nonce) + data)
		if strings.HasPrefix(hash, difficulty) {
			return nonce, hash
		}
		nonce++
	}
}

//NewBlock : creates a new block from a previous block and given data and difficulty
func NewBlock(block Block, data string, difficulty string) Block {
	t := time.Now().Unix()
	BlockID := block.BlockID
	BlockID++
	nonce, hash := computeHashWithProofOfWork(intToStr(BlockID)+intToStr(t)+data+block.Hash, difficulty)
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, nodeinfo.PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println(err)
	}
	return Block{BlockID, t, data, signature, nodeinfo.NodeID, block.Hash, hash, nonce}
}
