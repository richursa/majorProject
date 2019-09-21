package miner

import (
	"../db"
)

var peerlist = []string{"127.0.0.1"}

func mine() {
	client := db.GetClient()
	count := db.ReturnCount(client)

}
