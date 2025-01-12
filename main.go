package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents each block in the blockchain
type Block struct {
	Timestamp     int64
	Data          string
	PrevBlockHash string
	Hash          string
}

// Blockchain is a series of validated Blocks
type Blockchain struct {
	blocks []*Block
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() {
	data := fmt.Sprintf("%d%s%s", b.Timestamp, b.Data, b.PrevBlockHash)
	hash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hash[:])
}

// CreateBlock creates a new block using previous block's hash
func CreateBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.CalculateHash()
	return block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// CreateGenesisBlock creates the first block in the blockchain
func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", "")
}

// NewBlockchain creates a new Blockchain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

// IsValid checks if blockchain is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		previousBlock := bc.blocks[i-1]

		if currentBlock.PrevBlockHash != previousBlock.Hash {
			return false
		}

		var tempHash = currentBlock.Hash
		currentBlock.CalculateHash()
		if tempHash != currentBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	// Create new blockchain
	blockchain := NewBlockchain()

	// Add new blocks
	blockchain.AddBlock("First Block after Genesis")
	blockchain.AddBlock("Second Block after Genesis")
	blockchain.AddBlock("Third Block after Genesis")

	// Print all blocks
	for i, block := range blockchain.blocks {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

	// Validate blockchain
	fmt.Printf("Blockchain valid? %v\n", blockchain.IsValid())
}
