package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Cur_hash  []byte
	Prev_hash []byte
	Ts        int64
	Data      []byte
}

type Blockchain struct {
	blocks []*Block
}

// Setting hash
func (b *Block) SetHash() {
	ts := []byte(strconv.FormatInt(b.Ts, 10))
	headers := bytes.Join([][]byte{b.Prev_hash, b.Data, ts}, []byte{})
	cur_hash := sha256.Sum256(headers)
	b.Cur_hash = cur_hash[:]
}

// Adding block
func NewBlock(data string, prev_hash []byte) *Block {
	block := &Block{[]byte{}, prev_hash, time.Now().Unix(), []byte(data)}
	block.SetHash()
	return block
}

// Adding blockchain
func (bc *Blockchain) AddBlock(data string) {
	prev_block := bc.blocks[len(bc.blocks)-1]
	new_block := NewBlock(data, prev_block.Cur_hash)
	bc.blocks = append(bc.blocks, new_block)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Alice sends 1 BTC to Peter")
	bc.AddBlock("Peter sends 0.0004 BTC to Anna")

	for _, block := range bc.blocks {
		fmt.Printf("Previous hash: %x\n", block.Prev_hash)
		fmt.Printf("Current hash: %x\n", block.Cur_hash)
		fmt.Printf("Current data: %s\n", block.Data)
	}
}
