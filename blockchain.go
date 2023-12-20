package main

import (
	"time"
)

type Block struct {
	data map[string]interface {}
	hash string
	previousHash string
	timestamp time.Time
	pow int
}

type Blockchain struct {
	genesisBlock Block
	chain []Block
	difficulty int
}

func main() {
	// todo
}
