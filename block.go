package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	MerkleRoot    *MerkleTree
	PrevBlockHash []byte
	Hash          []byte
	//Nonce         int
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) HashTransactionsWithMerkle() *MerkleTree {
	var txs [][]byte

	for _, tx := range b.Transactions {
		txs = append(txs, tx.ID)
	}
	merkle_tree := CreateMerkleTree(txs)

	return merkle_tree
}

// Run performs a proof-of-work
func (block *Block) HashBlock() []byte {
	var hash [32]byte

	data := bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.HashTransactions(),
			IntToHex(block.Timestamp),
			block.MerkleRoot.Root.Data,
		},
		[]byte{},
	)

	hash = sha256.Sum256(data)
	fmt.Printf("\r%x", hash)

	fmt.Print("\n\n")

	return hash[:]
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, nil, prevBlockHash, []byte{}}
	//pow := NewProofOfWork(block)
	//nonce,
	//hash := pow.Run()
	block.MerkleRoot = block.HashTransactionsWithMerkle()
	block.Hash = block.HashBlock()

	//block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
