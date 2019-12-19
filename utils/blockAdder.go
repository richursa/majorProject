package main

import (
	"fmt"

	"../blockchain"
	"../db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	mongoClient := db.GetClient()
	lastBlockID := db.GetCount(mongoClient)
	var lastBlock blockchain.Block
	if lastBlockID == 0 {
		lastBlock.BlockID = 0
	} else {
		lastBlock = db.GetBlockFromDB(mongoClient, bson.M{"blockid": lastBlockID})
	}
	fmt.Println("last blockid is ", lastBlockID)
	fmt.Println("Enter the data for new block")
	var data string
	fmt.Scanln(&data)
	newBlock := blockchain.NewBlock(lastBlock, data)
	_ = db.InsertBlockIntoDB(mongoClient, newBlock)
}
