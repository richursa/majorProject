package main

import (
	"encoding/json"
	"fmt"

	"../miner"
)

func main() {
	count := miner.GetBlockCountFromPeer("http://localhost:8080")
	fmt.Println(count)
	block := miner.GetBlockFromPeer("http://localhost:8080", 1)
	j, _ := json.Marshal(block)
	fmt.Println(string(j))
}
