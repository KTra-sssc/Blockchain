package main

import (
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
	// Hash của block hiện tại
	Hash []byte
}

// Xây dựng cây Merkle từ danh sách các hash
func ConstructMerkleTree(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}
	if len(hashes) == 1 {
		return hashes[0]
	}

	var newHashes [][]byte

	for i := 0; i < len(hashes)-1; i += 2 {
		// Nối hai hash liên tiếp và tính toán hash mới
		concatenation := append(hashes[i], hashes[i+1]...)
		hash := sha256.Sum256(concatenation)
		newHashes = append(newHashes, hash[:])
	}

	// Nếu có số lượng hash lẻ, nhân bản hash cuối cùng
	if len(hashes)%2 != 0 {
		newHashes = append(newHashes, hashes[len(hashes)-1])
	}

	// Thực hiện đệ quy để xây dựng tiếp cây Merkle
	return ConstructMerkleTree(newHashes)
}

// Tính toán hash gốc của các giao dịch trong Block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte

	for _, tx := range b.Transactions {
		// Tính toán hash của từng giao dịch và thêm vào danh sách hash
		txHash := sha256.Sum256(tx.Data)
		txHashes = append(txHashes, txHash[:])
	}

	// Gọi hàm xây dựng cây Merkle từ danh sách hash
	return ConstructMerkleTree(txHashes)
}

// Tính toán và thiết lập hash cho Block
func (b *Block) SetHash() {
	// Kết hợp hash của block trước đó và hash của giao dịch
	headers := append(b.PrevBlockHash, b.HashTransactions()...)
	headers = append(headers, []byte(strconv.FormatInt(b.Timestamp, 10))...)

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
	}

	// Tính toán và thiết lập hash cho block
	block.SetHash()
	return block
}
