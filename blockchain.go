package main

// Blockchain đại diện cho chuỗi khối trong blockchain
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain tạo một chuỗi block mới với block khởi đầu (genesis block)
func NewBlockchain() *Blockchain {
	// Tạo Block khởi đầu không có giao dịch và hash trước đó
	genesisBlock := NewBlock([]*Transaction{}, []byte{})
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// Thêm một block mới vào chuỗi khối
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	// Lấy block trước đó từ chuỗi các block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	// Tạo một block mới với danh sách giao dịch và hash của block trước đó
	newBlock := NewBlock(transactions, prevBlock.Hash)
	// Thêm block mới vào chuỗi khối
	bc.Blocks = append(bc.Blocks, newBlock)
}
