package main

import (
	"./db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client := db.GetClient()
	block := db.ReturnBlockFromDB(client, bson.M{})
	println(block.Data)

}
