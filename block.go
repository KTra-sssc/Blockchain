package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Transaction struct {
	Data []byte
}

type Block struct {
	// Thời điểm tạo khối
	Timestamp int64
	// Danh sách các giao dịch trong khối
	Transactions []*Transaction
	// Hash của block trước đó
	PrevBlockHash []byte
	//merkle tree
	MerkleRoot *MerkleTree
	// Hash của block hiện tại
	Hash []byte
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Data)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) HashTransactionsWithMerkle() *MerkleTree {
	var txs [][]byte

	for _, tx := range b.Transactions {
		txs = append(txs, tx.Data)
	}
	merkle_tree := CreateMerkleTree(txs)

	return merkle_tree
}

// Tính toán và thiết lập hash cho Block
func (b *Block) SetHash() {
	// Kết hợp hash của block trước đó và hash của giao dịch
	headers := append(b.PrevBlockHash, b.HashTransactions()...)
	headers = append(headers, []byte(strconv.FormatInt(b.Timestamp, 10))...)
	headers = append(headers, b.MerkleRoot.Root.Data...)
	// Tính toán hash của toàn bộ block và thiết lập vào trường Hash
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock tạo một block mới với thời điểm, danh sách giao dịch và hash của block trước đó
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	// Khởi tạo block mới
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    nil,
	}
	//tạo merkle tree cho block
	block.MerkleRoot = block.HashTransactionsWithMerkle()
	// Tính toán và thiết lập hash cho block
	block.SetHash()
	return block
}
