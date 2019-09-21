package main

import "./blockchain"

func main() {
	richu := blockchain.Block{0, 2, "a", "a", "a", 5}
	secondBlock := blockchain.NewBlock(richu, "hello", "0000")
	println(secondBlock.BlockID)
}
