package main

import (
	"time"
	"fmt"
	"encoding/json"
	"crypto/sha256"
	"strconv"
	"strings"
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

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
		fmt.Println("mine hash: ", b.hash)
	}
	fmt.Println("done hash: ", b.hash)
}

func main() {
	// todo
}
