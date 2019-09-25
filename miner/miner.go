package miner

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"../blockchain"
	"../db"
)

var peerlist = []string{"http://127.0.0.1/"}
var client = db.GetClient()

//RequestBlock : request a block from peers and add block to chain
func RequestBlock() {
	client := db.GetClient()
	localCount := db.GetCount(client)
	for i := 0; i < len(peerlist); i++ {
		peerCount := GetBlockCountFromPeer(peerlist[i])
		if peerCount > localCount {
			for j := localCount + 1; j <= peerCount; j++ {
				block := GetBlockFromPeer(peerlist[i], j)
				//check if block is valid before approving it into db"
				db.InsertBlockIntoDB(client, block)
			}
		}
		if i == len(peerlist)-1 {
			time.Sleep(time.Second * 10)
			i = 0
		}
	}
}

//GetBlockCountFromPeer : get count of number of blocks a peer has
func GetBlockCountFromPeer(address string) int64 {
	address = address + "/api/getCount"
	resp, err := http.Get(address)
	if err != nil {
		log.Println("unable to send request to ", address, err)
		return 0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("unable to parse response from ", address, err)
		return 0
	}
	count, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		log.Println("cannot convert count to integer from ", address, err)
		return 0
	}
	return count
}

//GetBlockFromPeer : get a block with given blockid from peer
func GetBlockFromPeer(address string, blockID int64) blockchain.Block {
	address = address + "/api/getBlock/"
	address = address + strconv.FormatInt(blockID, 10)
	resp, err := http.Get(address)
	if err != nil {
		log.Println("unable to send request to ", address, err)
		return blockchain.Block{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("unable to parse response from ", address, err)
		return blockchain.Block{}
	}
	block := blockchain.Block{}
	json.Unmarshal(body, &block)
	return block
}
