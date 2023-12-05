package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

type CLI struct {
	bc *Blockchain
}

func NewCLI(bc *Blockchain) *CLI {
	return &CLI{bc: bc}
}

// PrintUsage in ra thông tin sử dụng dòng lệnh
func PrintUsage() {
	fmt.Println("Cac lenh:")
	fmt.Println("  -create <blockchain name> : Tao blockchain moi.")
	fmt.Println("  -add <value> -blockchain <blockchain name>: Them giao dich vao blockchain cu the.")
	//fmt.Println("  -mine <blockchain name>: Khai thac mot blockchain.")
	fmt.Println("  -print <blockchain name>: In thong tin blockchain.")
	fmt.Println("  -exit - Thoat")
}

func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Can cung cap command-line arguments.")
		PrintUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "-create":
		cli.createBlockchain()
	case "-add":
		cli.addTransaction()
	// case "-mine":
	// 	cli.mineBlock()
	case "-print":
		cli.printBlockchain()
	case "-exit":
		os.Exit(0)
	default:
		fmt.Println("Lenh khong hop le:", command)
		PrintUsage()
		os.Exit(1)
	}
}

// OpenDB mở một cơ sở dữ liệu BoltDB với tên blockchain được cung cấp
func OpenDB(blockchainName string) (*bolt.DB, error) {
	db, err := bolt.Open(blockchainName+".db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// BlockchainExists kiểm tra xem một blockchain với tên đã cho có tồn tại hay không
func BlockchainExists(blockchainName string) bool {
	// Mở cơ sở dữ liệu với quyền chỉ đọc
	db, err := bolt.Open(blockchainName+".db", 0400, nil)
	if err != nil {
		return false
	}
	defer db.Close()

	// Kiểm tra xem có "Blockchain" bucket tồn tại hay không
	exists := false
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Blockchain"))
		if b != nil {
			exists = true
		}
		return nil
	})

	return exists
}

// Tạo một blockchain mới
func (cli *CLI) createBlockchain() {
	if len(os.Args) < 3 {
		fmt.Println("Can nhap ten blockchain.")
		os.Exit(1)
	}

	blockchainName := os.Args[2]
	db, err := OpenDB(blockchainName)
	if err != nil {
		fmt.Println("Loi: ", err)
		os.Exit(1)
	}
	defer db.Close()

	cli.bc = NewBlockchain(db)
	fmt.Println("Blockchain", blockchainName, " tao thanh cong.")
}

// Thêm một giao dịch vào khối hiện tại
func (cli *CLI) addTransaction() {
	if os.Args[1] != "-add" || os.Args[3] != "-blockchain" {
		fmt.Println("Lenh khong ton tai.")
		os.Exit(1)
	}

	value := os.Args[2]
	blockchainName := os.Args[4]

	db, err := OpenDB(blockchainName)
	if err != nil {
		fmt.Println("Loi: ", err)
		os.Exit(1)
	}
	defer db.Close()

	cli.bc = NewBlockchain(db)

	// Tạo một giao dịch mới
	transaction := &Transaction{Data: []byte(value)}

	// Thêm giao dịch vào khối hiện tại
	cli.bc.AddBlock([]*Transaction{transaction})

	fmt.Println("Giao dich duoc them vao khoi hien tai.")
}

// mineBlock mines a new block
// func (cli *CLI) mineBlock() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("Please provide a blockchain name.")
// 		os.Exit(1)
// 	}

// 	blockchainName := os.Args[2]

// 	db, err := OpenDB(blockchainName)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		os.Exit(1)
// 	}
// 	defer db.Close()

// 	cli.bc = NewBlockchain(db)

// 	mineBlock(cli.bc)
// }

// In ra trạng thái hiện tại của blockchain
func (cli *CLI) printBlockchain() {
	if len(os.Args) < 3 {
		fmt.Println("Can nhap ten blockchain.")
		os.Exit(1)
	}

	blockchainName := os.Args[2]

	db, err := OpenDB(blockchainName)
	if err != nil {
		fmt.Println("Loi: ", err)
		os.Exit(1)
	}
	defer db.Close()

	cli.bc = NewBlockchain(db)

	printBlockchain(cli.bc)
}
