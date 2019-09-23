package miner

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../blockchain"
)

var peerlist = []string{"127.0.0.1"}

/*func RequestBlock() {
	client := db.GetClient()
	localCount := db.GetCount(client)
}*/
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
