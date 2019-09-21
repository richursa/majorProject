package main

import (
	"./blockchain"
	"./db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client := db.GetClient()
	block := db.ReturnBlockFromDB(client, bson.M{})
	println(block.Data)
	db.InsertBlockIntoDB(client, blockchain.NewBlock(block, "hello", "0000"))
	block = db.ReturnBlockFromDB(client, bson.M{"data": "hello"})
	println(block.Data)
	println(db.ReturnCount(client))
}
