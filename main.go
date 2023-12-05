package main

import (
	"bufio"
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
	fmt.Println("  - mine: Tinh merkle tree va hash cua block, dong block va tao block moi.")
	fmt.Println("  - print: In toan bo chuoi khoi.")
	fmt.Println("  - exit: Thoat khoi CLI.")
}
