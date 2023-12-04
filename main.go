package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

var bc *Blockchain

func main() {
	// Tạo một chuỗi block mới (tạo blockchain)
	bc := NewBlockchain()

	fmt.Println("Blockchain CLI. Go 'help' de xem danh sach cac lenh.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "help":
			printHelp()
		case "add":
			addTransaction(bc)
		case "mine":
			mineBlock(bc)
		case "print":
			printBlockchain(bc)
		case "exit":
			fmt.Println("Thoat...")
			os.Exit(0)
		default:
			fmt.Println("Lenh khong ton tai. Go 'help' de xem danh sach cac lenh.")
		}
	}
}

func printHelp() {
	fmt.Println("Cac lenh:")
	fmt.Println("  - add: Them giao dich moi vao khoi hien tai.")
	fmt.Println("  - mine: Them khoi vao chuoi khoi.")
	fmt.Println("  - print: In toan bo chuoi khoi.")
	fmt.Println("  - exit: Thoat khoi CLI.")
}

func addTransaction(bc *Blockchain) {
	fmt.Print("Nhap du lieu giao dich: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	tx := &Transaction{Data: []byte(data)}
	currentBlock := bc.Blocks[len(bc.Blocks)-1]
	currentBlock.Transactions = append(currentBlock.Transactions, tx)
	fmt.Println("Giao dich duoc them vao khoi hien tai.")
}

func mineBlock(bc *Blockchain) {
	currentBlock := bc.Blocks[len(bc.Blocks)-1]
	// Tính toán lại hash sau khi thêm giao dịch
	currentBlock.SetHash()
	newBlock := NewBlock([]*Transaction{}, currentBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
	fmt.Println("Khoi da duoc them vao chuoi khoi.")
}

func printBlockchain(bc *Blockchain) {
	fmt.Println("Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Block %x\n", block.Hash)
		for _, tx := range block.Transactions {
			fmt.Printf("  Giao dich %x\n", sha256.Sum256(tx.Data))
		}
	}
}
