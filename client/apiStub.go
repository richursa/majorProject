package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func getBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("request received ")
	blockID, err := strconv.ParseInt(vars["blockID"], 10, 64)
	if err != nil {
		log.Println("couldnt covert blockID to int", err)
		return
	}
	blockID++
	block := Block{1, 1212, "this is a sample block", []uint8{2}, "cusatNode", "dfs3d3j8ggjs7axsa7dvnza39kaazvbwma8hd72nas", "as82s8alzu7ehvzdneueud73jaks3kskmak22ks3he7dfa", 1992}
	j, err := json.Marshal(block)
	if err != nil {
		log.Fatalln("unable to parse json ", err)
	}
	w.Write([]byte(j))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/getBlock/{blockID}", getBlock)
	http.ListenAndServe(":8080", r)
}
