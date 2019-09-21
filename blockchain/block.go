package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	BlockID int64  //stores block number
	Time    int64  //to store time
	Data    string //transactions/data which is to be stored in a block
	Prev    string
	Hash    string
	Nonce   int64
}

func calcHash(data string) string {
	hashed := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashed[:])
}
func intToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
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

func NewBlock(block Block, data string, difficulty string) Block {
	t := time.Now().Unix()
	BlockID := block.BlockID
	BlockID++
	nonce, hash := computeHashWithProofOfWork(intToStr(BlockID)+intToStr(t)+data+block.Hash, difficulty)
	return Block{BlockID, t, data, block.Hash, hash, nonce}
}
