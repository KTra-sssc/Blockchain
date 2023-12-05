package main

import (
	"crypto/sha256"
	"fmt"
)

func addTransaction(bc *Blockchain, data string) {
	tx := &Transaction{Data: []byte(data)}
	currentBlock := bc.Blocks[len(bc.Blocks)-1]
	currentBlock.Transactions = append(currentBlock.Transactions, tx)
	fmt.Println("Giao dich duoc them vao khoi hien tai.")
	currentBlock.SetHash()

}

func mineBlock(bc *Blockchain) {
	currentBlock := bc.Blocks[len(bc.Blocks)-1]
	currentBlock.MerkleRoot = currentBlock.HashTransactionsWithMerkle()
	// Tính toán lại hash sau khi thêm giao dịch
	currentBlock.SetHash()
	newBlock := NewBlock([]*Transaction{}, currentBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
	fmt.Printf("Khoi %d da duoc dong va khoi moi da tao.\n", len(bc.Blocks)-2)
}

func printBlockchain(bc *Blockchain) {
	fmt.Println("Blockchain:")
	for i, block := range bc.Blocks {
		fmt.Printf("#%d\n", i)
		fmt.Printf("  Prev: %x\n", block.PrevBlockHash)
		fmt.Printf("  Block Hash: %x\n", block.Hash)
		for _, tx := range block.Transactions {
			fmt.Printf("  Giao dich %x\n", sha256.Sum256(tx.Data))
		}
		fmt.Printf("  Merkle Root: %x\n", block.MerkleRoot.Root.Data)
	}
}
