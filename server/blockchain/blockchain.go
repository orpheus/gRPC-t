package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash string
	PrevBlockHash string
	Data string
}

type Blockchain struct {
	//what would it be like if this didn't have a pointer?
	//if these point all to Block, how could the same struct get multiplied so many times?
	//where does a struct get duplicated (for many blocks to be)
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, prevBlockHash string) *Block {
	//block = new(Block) ?
	block := &Block{
		Data: data,
		PrevBlockHash: prevBlockHash,
	}

	//is 'block' the block that gets paramterized to b *Block in setHash?
	block.setHash()

	return block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blockchain {
	//why would you want a new blockchain? why not a copy?
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

